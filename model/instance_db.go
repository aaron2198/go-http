package model

import "time"

type InstanceDb struct {
	ID        uint64 `gorm:"PrimaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Hostname  string `gorm:"notNull"`
	Port      int    `gorm:"notNull"`
	Username  string `gorm:"notNull"`
	Password  string `gorm:"notNull"`
}
