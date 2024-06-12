package models

import "time"

type Company struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Name        string
	Image       string
	Slug        string
	Description string
	Location    string
	Website     string
	GitHub      string `gorm:"column:github"`
	LinkedIn    string `gorm:"column:linkedin"`
}
