package model

import "time"

type Community struct {
	ID           uint `gorm:"PrimaryKey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Subdomain    string `gorm:"index;unique"`
	DatabaseID   uint
	Database     InstanceDb
	DatabaseUser string `gorm:"notNull"`
	DatabasePass string `gorm:"notNull"`
}
