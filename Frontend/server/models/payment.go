package models

import "gorm.io/gorm"

type Payment struct {
	gorm.Model

	IsPaid bool

	CartID uint
	Cart   Cart

	CustomerID uint
	Customer   Customer
}
