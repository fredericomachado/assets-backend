package models

import (
	"gorm.io/gorm"
)

type Asset struct {
	gorm.Model
	Name        string `json:"nome"`
	Ticker      string
	LocationID  int
	Location    Location
	CategoryID  int
	Category    Category
	Allocations []Allocation
}
