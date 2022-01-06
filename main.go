package main

import (
	"bot/market"
	"context"
	"flag"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/log"
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
	cycle   = flag.Int("cycle", 60, "Scan at an interval time(minute)")
)

func main() {
	// Parse the flags and set up the logger to print everything requested
	flag.Parse()
	log.Root().SetHandler(log.LvlFilterHandler(log.Lvl(*logFlag), log.StreamHandler(os.Stderr, log.TerminalFormat(true))))

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

func loop(ctx context.Context) {
	var (
		scope      event.SubscriptionScope
		createSink = make(chan *market.VenusMarketCreateOrder, 16)
		delSink    = make(chan *market.VenusMarketCancelOrder, 16)
		retry      = time.Minute * 5
		errMsg     = make(chan error, 2)
	)
	defer scope.Close()

	scope.Track(event.ResubscribeErr(retry, func(ctx context.Context, err error) (event.Subscription, error) {
		if err != nil {
			errMsg <- err
		}
		return cfg.MarketSession.Contract.WatchCreateOrder(&bind.WatchOpts{Context: ctx, Start: &cfg.StartNumber}, createSink)
	}))

	scope.Track(event.ResubscribeErr(retry, func(ctx context.Context, err error) (event.Subscription, error) {
		if err != nil {
			errMsg <- err
		}
		return cfg.MarketSession.Contract.WatchCancelOrder(&bind.WatchOpts{Context: ctx, Start: &cfg.StartNumber}, delSink)
	}))

	tick := time.NewTicker(time.Minute * 10)
	defer tick.Stop()

	for {
		select {
		case sink := <-createSink:
			checkErr("failed to insert order", InsertOrder(&VenusMarket{
				Number: sink.Raw.BlockNumber,
				TxHash: sink.Raw.TxHash.String(),

				Id:      sink.Oid.Int64(),
				Address: sink.User.String(),
				Price:   sink.Price.String(),
			}))
		case sink := <-delSink:
			checkErr("failed to delete order", DelOrder(sink.Oid.Int64()))

		case <-tick.C:
			checkErr("failed to scan orders", scanOrders())

		case err := <-errMsg:
			log.Error("failed to watch order", "error", err)

		case <-ctx.Done():
			return
		}
	}
}

func loop2(ctx context.Context) {
	var (
		tick     = time.NewTicker(time.Minute * time.Duration(*cycle))
		contract = cfg.MarketSession
	)
	defer tick.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-tick.C:
			id, err := contract.OrderID()
			if err != nil {
				log.Error("Failed to get orderId", "err", err)
				continue
			}
			orderId := id.Int64()
			for i := 10000; i <= int(orderId); i++ {
				oid := big.NewInt(int64(i))
				order, err := contract.Orders(oid)
				if err != nil {
					log.Error("failed to get order", "err", err)
				}
				stepTime, endTime, curTime := order.StepTime.Int64(), order.EndTime.Int64(), time.Now().Unix()
				if order.Status && endTime > curTime && stepTime < curTime {
					_, err := contract.TrigOrder(oid)
					if err != nil {
						log.Error("failed to trigger order", "orderId", i, "err", err)
					}
				}
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
