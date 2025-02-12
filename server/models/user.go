package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName  string       `json:"first_name" gorm:"size:100;not null"`
	LastName   string       `json:"last_name" gorm:"size:100;not null"`
	Email      string       `json:"email" gorm:"size:100;unique;not null"`
	Password   string       `json:"-" gorm:"not null"`
	RealEstate []RealEstate `json:"real_estate" gorm:"foreighKey:UserID"`
}
