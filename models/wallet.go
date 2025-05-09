package models

import "gorm.io/gorm"

type Wallet struct {
	gorm.Model
	Name                   string `gorm:"not null" json:"name"`
	UserID                 int
	User                   User
	Allocations            []Allocation
	TargetMacroAllocations []TargetMacroAllocation
	TargetAllocations      []TargetAllocation
}
