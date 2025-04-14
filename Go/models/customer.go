package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string
	CartItems []Cart `gorm:"foreignKey:CustomerID"`
}
