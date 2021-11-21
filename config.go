package main

import (
	"bot/market"
	"bytes"
	"encoding/json"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"os"
)

type Config struct {
	ChainId       *big.Int                   `json:"chain_id"`
	Endpoint      string                     `json:"endpoint"`
	Keystore      string                     `json:"keystore"`
	MarketAddress common.Address             `json:"market_address"`
	DBurl         string                     `json:"db_url"`
	MarketSession *market.VenusMarketSession `json:"-"`
}

func NewConfig(path string, password string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	cfg := &Config{}
	if err := json.Unmarshal(data, cfg); err != nil {
		return nil, err
	}
	store, err := os.ReadFile(cfg.Keystore)
	if err != nil {
		utils.Fatalf("keystore read failed", err)
	}
	account, err := bind.NewTransactorWithChainID(bytes.NewReader(store), password, cfg.ChainId)
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
