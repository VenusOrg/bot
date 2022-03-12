package store

import (
	"github.com/ethereum/go-ethereum/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

const (
	Creater = iota
	Cancel
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
	Id   int64 `gorm:"primaryKey"`
	Type int
	// block message
	Number   uint64 `gorm:"index"`
	TxHash   string `gorm:"index"`
	User     string `gorm:"index"`
	TokenIn  string
	TokenOut string
	Price    string
}

func InsertOrder(order *VenusMarket) error {
	rows, err := db.Model(&VenusMarket{}).Where("id = ?", order.Id).Rows()
	if err == nil && rows.Next() {
		return nil
	}
	return db.Create(order).Error
}

func DelOrder(oid int64) error {
	return db.Where("oid = ?", oid).Delete(&VenusMarket{}).Error
}

func ScanOrders() ([]int64, error) {
	var ids []int64
	return ids, db.Model(&VenusMarket{}).Pluck("id", &ids).Error
}

func UpdateOrder(id int64) error {
	return db.Model(&VenusMarket{}).Where("id = ?", id).Update("type", gorm.Expr("1")).Error
}

func GetOrder(id int64) (*VenusMarket, error) {
	res := &VenusMarket{}
	if err := db.First(res, id).Error; err != nil {
		return nil, err
	}
	return res, nil
}
