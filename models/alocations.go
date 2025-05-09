package models

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Allocation struct {
	gorm.Model
	UserID        int `json:"user_id"`
	User          User
	WalletID      int `json:"wallet_id"`
	Wallet        Wallet
	AssetID       int `json:"asset_id"`
	Asset         Asset
	Quantity      int `json:"quantity"`
	PurchasePrice decimal.Decimal
	CurrentPrice  decimal.Decimal
	MidPrice      decimal.Decimal
}
