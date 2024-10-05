package models

import "time"

type Products struct {
	ID          uint
	Name        string
	Description string
	Price       float64
	Count       uint
	CategoryID  uint
	Category    Category
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
