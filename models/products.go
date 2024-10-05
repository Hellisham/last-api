package models

import "time"

type Products struct {
	ID          uint
	Name        string
	Description string
	Price       float64
	Count       uint
	CategoryID  uint
	Category    Category `gorm:"foreignKey:CategoryID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
