package models

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

// Product is full model, used by gorm
type Product struct {
	gorm.Model
	Title    string          `gorm:"size:255"`
	Price    decimal.Decimal `json:"price" sql:"type:decimal(20,8);"`
	StatusID uint
	Status   Status `json:"status"` // used to populate obj for create
}

// ProductCreateInput for capturing user input for add
// no gorm info, need json
// leaving out ID, 3 date fields
type ProductCreateInput struct {
	Title  string          `json:"title" binding:"required"`
	Price  decimal.Decimal `json:"price" binding:"required"`
	Status Status          `json:"status" binding:"required"` // input as status.status (string, not id, that needs to be converted)
}

// ProductOutput for clean/santitized version back to user
// it should be stripped of private data
// keep ID, drop 3 date fields
type ProductOutput struct {
	ID    uint
	Title string
	Price decimal.Decimal
	// Status string
}

// ID          uint   `json:"id"`
// Status      string `json:"status"`
// Description string `json:"description"`
// Active      bool   `json:"active"`
