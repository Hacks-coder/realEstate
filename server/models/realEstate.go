package models

import "gorm.io/gorm"

type RealEstate struct {
	gorm.Model
	Category string  `json:"category" gorm:"size:200;not null"`
	Amount   float64 `json:"amount" gorm:"size:200;not null"`
	Location string  `json:"location" gorm:"size:200;not null"`
	Types    string  `json:"types" gorm:"size:200;not null"`
	Photo    string  `json:"photo" gorm:"size:200;not null"`
	UserID   uint    `json:"user_id" gorm:"not null"`
}
