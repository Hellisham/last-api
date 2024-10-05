package models

import "time"

type Category struct {
	ID          uint
	Name        string
	Description string
	Created_at  time.Time
	Updated_at  time.Time

	Products []Products `json:"products" gorm:"foreignKey:CategoryID"`
}
