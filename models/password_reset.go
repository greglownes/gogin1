package models

import "gorm.io/gorm"

type PasswordReset struct {
	gorm.Model
	UserID uint   `gorm:"not null"`
	Token  string `gorm:"not null;unique_index"`
}
