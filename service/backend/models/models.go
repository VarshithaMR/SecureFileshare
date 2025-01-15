package models

import "time"

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Role      string    `json:"role"`
	MFASecret string    `json:"mfa_secret"`
	CreatedAt time.Time `json:"created_at"`
}