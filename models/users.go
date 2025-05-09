package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name        string       `json:"nome"`
	Wallets     []Wallet     `json:"wallets"`
	Allocations []Allocation `json:"allocations"`
}
