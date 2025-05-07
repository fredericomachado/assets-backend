package models

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Asset struct {
	gorm.Model
	UserID        int
	Name          string `json:"nome"`
	Ticker        string
	Quantity      int
	PurchasePrice decimal.Decimal
	CurrentPrice  decimal.Decimal
	MidPrice      decimal.Decimal
	Percentage    decimal.Decimal
}
