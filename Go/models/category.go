package models

type Category struct {
	ID       uint      `gorm:"primarykey"`
	Name     string    `json:"name"`
	Products []Product `json:"products,omitempty"`
}
