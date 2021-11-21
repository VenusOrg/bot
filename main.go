package main

import (
	"bot/market"
	"context"
	"flag"
	"github.com/ethereum/go-ethereum/cmd/utils"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	cfg      *Config
	path     = flag.String("config", "./config.json", "config file path")
	password = flag.String("password", "123456", "keystore password")
)

func main() {
	var err error
	cfg, err = NewConfig(*path, *password)
	if err != nil {
		utils.Fatalf("config load failed", err)
	}

	ctx := context.Background()
	go loop(ctx)

	watcher()
}

func watcher() {
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(sigc)
}

func loop(ctx context.Context) {
	createSink := make(chan *market.VenusMarketCreateOrder)
	createSub, _ := cfg.MarketSession.Contract.WatchCreateOrder(nil, createSink)
	defer createSub.Unsubscribe()

	delSink := make(chan *market.VenusMarketCancelOrder)
	cfg.MarketSession.Contract.WatchCancelOrder(nil, delSink)
	delSub, _ := cfg.MarketSession.Contract.WatchCancelOrder(nil, delSink)
	defer delSub.Unsubscribe()

	tick := time.NewTicker(time.Hour * 1)
	defer tick.Stop()

	for {
		select {
		case sink := <-createSink:
			InsetOrder(&VenusMarket{
				Oid:     sink.Oid.Uint64(),
				Address: sink.User.String(),
				Price:   sink.Price.String(),
				Cycle:   sink.Cycle.Uint64(),
				EndTime: sink.EndTime.Uint64(),
			})
		case sink := <-delSink:
			DelOrder(sink.Oid.Uint64())

		case <-tick.C:
			scanOrders()

		case <-ctx.Done():
			return
		}
	}
}

func scanOrders() {
}
