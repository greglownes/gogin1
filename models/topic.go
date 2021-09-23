package models

import "gorm.io/gorm"

type Topic struct {
	gorm.Model
	Title string `gorm:"size:255"`
}
