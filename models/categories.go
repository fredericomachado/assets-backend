package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name              string  `gorm:"not null" json:"name"`
	Assets            []Asset `json:"assets"`
	TargetAllocations []TargetAllocation
}
