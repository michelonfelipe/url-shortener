package models

import "gorm.io/gorm"

type Url struct {
	gorm.Model
	Original  string `gorm:"uniqueIndex"`
	Shortened string `gorm:"uniqueIndex"`
}
