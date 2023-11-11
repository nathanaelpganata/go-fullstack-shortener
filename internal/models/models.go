package models

import (
	"time"

	"gorm.io/gorm"
)

type Link struct {
	gorm.Model
	OriginalURL string    `gorm:"not null" form:"originalUrl"`
	ShortURL    string    `gorm:"unique; not null" form:"shortUrl"`
	ExpiresAt   time.Time `gorm:"default:null"`
	Clicks      int       `gorm:"default:0"`
	RequestIP   string    `gorm:"default:null"`
}
