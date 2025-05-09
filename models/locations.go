package models

import "gorm.io/gorm"

type Location struct {
	gorm.Model
	Name                   string `gorm:"not null" json:"name"`
	Assets                 []Asset
	TargetMacroAllocations []TargetMacroAllocation
}
