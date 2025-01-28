package models

import "time"

type User struct {
	ID              uint   `gorm:"primaryKey"`
	Email           string `gorm:"unique;not null"`
	PasswordHash    string `gorm:"not null"`
	Phone           string
	IsEmailVerified bool      `gorm:"default:false"`
	CreatedAt       time.Time `gorm:"autoCreateTime"`
}
