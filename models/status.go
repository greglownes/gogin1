package models

import "gorm.io/gorm"

type Status struct {
	gorm.Model
	Status      string `json:"status" gorm:"size:25;index:idx_status,unique"`
	Description string `gorm:"size:255"`
	Active      bool   `gorm:"type:bool;default:true"`
	// Products    []Product
}
