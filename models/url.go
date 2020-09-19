package models

import "gorm.io/gorm"

type Url struct {
	gorm.Model
	Original  string `gorm:"index"`
	Shortened string `gorm:"index"`
}
