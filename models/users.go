package models

import (
	"time"
	_ "time"
)

type User struct {
	ID        uint
	Name      string `gorm:"unique"`
	Email     string `gorm:"unique"`
	Password  string `gorm:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
