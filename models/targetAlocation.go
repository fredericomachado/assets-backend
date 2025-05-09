package models

import "gorm.io/gorm"

type TargetAllocation struct {
	gorm.Model
	WalletID           int `json:"wallet_id"`
	Wallet             Wallet
	CategoryID         int `json:"category_id"`
	Category           Category
	CategoryPercentage int `json:"percentage"`
}
