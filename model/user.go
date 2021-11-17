package model

import (
	"time"
)

type User struct {
	ID           uint `gorm:"PrimaryKey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Email        string `gorm:"index;unique"`
	FirstName    string `gorm:"notNull"`
	LastName     string `gorm:"notNull"`
	Password     string `gorm:"notNull"`
	DatabaseID   uint
	DatabaseUser string `gorm:"notNull"`
	DatabasePass string `gorm:"notNull"`
}
