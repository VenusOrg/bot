package main

import (
	"bot/config"
	"bot/market"
	"bot/store"
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/log"
	"math/big"
	"time"
)

func loop2(ctx context.Context) {
	tick := time.NewTicker(time.Minute * time.Duration(*cycle))
	defer tick.Stop()

	ctCh := make(chan *market.VenusMarketCreateOrder, 128)
	sb, err := cfg.Market.Contract.WatchCreateOrder(&bind.WatchOpts{Context: ctx}, ctCh)
	if err != nil {
		log.Crit("failed to subscribe createOrder", "err", err)
	}
	defer sb.Unsubscribe()

	cancelCh := make(chan *market.VenusMarketCancelOrder, 128)
	cancelSb, err := cfg.Market.Contract.WatchCancelOrder(&bind.WatchOpts{Context: ctx}, cancelCh)
	if err != nil {
		log.Crit("failed to subscribe cancel order", "err", err)
	}
	defer cancelSb.Unsubscribe()

	for {
		select {
		case res := <-ctCh:
			if err := store.InsertOrder(&store.VenusMarket{
				Id:       res.Oid.Int64(),
				Type:     store.Creater,
				User:     res.User.String(),
				TokenIn:  res.TokenIn.String(),
				TokenOut: res.TokenOut.String(),
				Price:    res.Price.String(),
				Number:   res.Raw.BlockNumber,
				TxHash:   res.Raw.TxHash.String(),
			}); err != nil {
				log.Error("failed to insert into db", "err", err)
			}
		case res := <-cancelCh:
			if err := store.UpdateOrder(res.Oid.Int64()); err != nil {
				log.Error("failed to update order status", "err", err)
			}

		case <-tick.C:
			scanOrders2()
		}
	}
}

func scanOrders2() {
	contract := cfg.Market

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
	ids, err := store.ScanOrders()
	if err != nil {
		return err
	}
	for _, id := range ids {
		config.CheckErr("failed to trig order", trigOrder(id))
	}
	return nil
}

func trigOrder(id int64) error {
	start := time.Now().Unix()
	tmp, err := cfg.Market.Orders(big.NewInt(id))
	if err != nil {
		return err
	}
	if !tmp.Status {
		return store.DelOrder(id)
	}
	if tmp.EndTime.Int64() > start || tmp.StepTime.Int64() > start {
		return nil
	}
	_, err = cfg.Market.TrigOrder(big.NewInt(id))
	if err == nil {
		config.CheckErr("failed to update trig times", store.UpdateOrder(id))
	}
	return err
}
