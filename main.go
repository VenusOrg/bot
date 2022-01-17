package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/log"
	"io"
	"math/big"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	cfg *Config

	path    = flag.String("config", "./config.json", "config file path")
	logFlag = flag.Int("loglevel", 3, "Log level to use for proxyServer")
	logFile = flag.String("logfile", "", "Tells the bot where to write log entries")
	cycle   = flag.Int("cycle", 60, "Scan at an interval time(minute)")
)

func main() {
	// Parse the flags and set up the logger to print everything requested
	flag.Parse()
	stream := log.StreamHandler(os.Stderr, log.TerminalFormat(true))
	if len(*logFile) != 0 {
		fp, err := os.OpenFile(*logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
		if err != nil {
			fmt.Println("Failed to open log file", "err", err)
			os.Exit(1)
		}
		stream = log.StreamHandler(io.Writer(fp), log.TerminalFormat(true))
	}
	glogger := log.NewGlogHandler(stream)
	glogger.Verbosity(log.Lvl(*logFlag))
	log.Root().SetHandler(glogger)

	var err error
	cfg, err = NewConfig(*path)
	if err != nil {
		utils.Fatalf("config load failed", err)
	}

	ctx := context.Background()
	defer ctx.Done()

	go loop2(ctx)
	watcher()
}

func watcher() {
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(sigc)
	<-sigc
}

func loop2(ctx context.Context) {
	tick := time.NewTicker(time.Minute * time.Duration(*cycle))
	defer tick.Stop()

	// The first time scan orders.
	scanOrders2()
	for {
		select {
		case <-ctx.Done():
			return
		case <-tick.C:
			scanOrders2()
		}
	}
}

func scanOrders2() {
	contract := cfg.MarketSession

	id, err := contract.OrderID()
	if err != nil {
		log.Error("Failed to get orderId", "err", err)
		return
	}
	log.Info("Get the latest orderId", "orderId", id.Int64())

	orderId := id.Int64()
	for i := 10000; i <= int(orderId); i++ {
		oid := big.NewInt(int64(i))
		order, err := contract.Orders(oid)
		if err != nil {
			log.Error("failed to get order", "err", err)
		}
		log.Info("Get order detail", "order.address", order.Owner.String(), "order.status", order.Status)

		stepTime, endTime, curTime := order.StepTime.Int64(), order.EndTime.Int64(), time.Now().Unix()
		if order.Status && endTime > curTime && stepTime < curTime {
			_, err := contract.TrigOrder(oid)
			if err != nil {
				log.Error("failed to trigger order", "orderId", i, "err", err)
			}
		}
	}
}

func scanOrders() error {
	ids, err := ScanOrders()
	if err != nil {
		return err
	}
	for _, id := range ids {
		checkErr("failed to trig order", trigOrder(id))
	}
	return nil
}

func trigOrder(id int64) error {
	start := time.Now().Unix()
	tmp, err := cfg.MarketSession.Orders(big.NewInt(id))
	if err != nil {
		return err
	}
	if !tmp.Status {
		return DelOrder(id)
	}
	if tmp.EndTime.Int64() > start || tmp.StepTime.Int64() > start {
		return nil
	}
	_, err = cfg.MarketSession.TrigOrder(big.NewInt(id))
	if err == nil {
		checkErr("failed to update trig times", UpdateOrder(id))
	}
	return err
}
