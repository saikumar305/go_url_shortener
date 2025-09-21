package models

import "time"


type URL struct {
	ID        uint   `gorm:"primaryKey"`
	Code      string `gorm:"unique;not null"`
	Original  string `gorm:"not null"`
	ShortUrl  string `gorm:"unique;not null"`
	CreatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`
}