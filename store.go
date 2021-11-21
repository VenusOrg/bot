package main

import (
	"github.com/ethereum/go-ethereum/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func init() {
	var err error
	db, err = gorm.Open(sqlite.Open("venus_market.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	if !db.Migrator().HasTable(&VenusMarket{}) {
		log.Info("create venus_market table")
		db.AutoMigrate(&VenusMarket{})
	}
}

type VenusMarket struct {
	Oid      uint64
	Address  string
	Price    string
	Cycle    uint64
	EndTime  uint64
	ScanTime uint64
}

func InsetOrder(market *VenusMarket) error {
	return db.Create(market).Error
}

func DelOrder(oid uint64) error {
	return db.Where("oid = ?", oid).Delete(&VenusMarket{}).Error
}

func UpdateScanTime(oid uint64, scanTime uint64) error {
	return db.Model(&VenusMarket{}).Updates(VenusMarket{Oid: oid, ScanTime: scanTime}).Error
}
