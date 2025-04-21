package models

import "gorm.io/gorm"

type Payment struct {
	gorm.Model

	IsPaid bool
	Amount uint
}
