package main

import (
	"bot/market"
	"encoding/json"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"math/big"
	"os"
)

type Config struct {
	StartNumber   uint64                     `json:"start_number"`
	ChainId       *big.Int                   `json:"chain_id"`
	Endpoint      string                     `json:"endpoint"`
	TrigPrivate   string                     `json:"trig_private"`
	MarketAddress common.Address             `json:"market_address"`
	DBurl         string                     `json:"db_url"`
	MarketSession *market.VenusMarketSession `json:"-"`
}

func NewConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	cfg := &Config{}
	if err := json.Unmarshal(data, cfg); err != nil {
		return nil, err
	}
	priv, err := crypto.ToECDSA(common.FromHex(cfg.TrigPrivate))
	if err != nil {
		return nil, err
	}
	account, err := bind.NewKeyedTransactorWithChainID(priv, cfg.ChainId)
	if err != nil {
		utils.Fatalf("failed to create account", err)
	}

	client, err := ethclient.Dial(cfg.Endpoint)
	if err != nil {
		utils.Fatalf("failed to connect to endpoint")
	}

	mkt, _ := market.NewVenusMarket(cfg.MarketAddress, client)

	cfg.MarketSession = &market.VenusMarketSession{
		Contract:     mkt,
		TransactOpts: *account,
	}

	return cfg, nil
}

func checkErr(msg string, err error) {
	if err != nil {
		log.Error(msg, "error", err)
	}
}
