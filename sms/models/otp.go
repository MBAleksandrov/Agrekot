package models

import "time"

type OTP struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Phone     string    `json:"phone"`
	Code      string    `json:"code"`
	ExpiresAt time.Time `json:"expires_at"`
	IsUsed    bool      `json:"is_used"`
}
