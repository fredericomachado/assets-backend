package models

import "gorm.io/gorm"

type TargetMacroAllocation struct {
	gorm.Model
	LocationID int `json:"location_id"`
	Location   Location
	WalletID   int `json:"wallet_id"`
	Wallet     Wallet
	Percentage int `json:"percentage"`
}
