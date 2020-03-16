package user

import "time"

type User struct {
	ID              uint       `json:"id,omitempty"`
	Name            string     `json:"name" validate:"required"`
	Email           string     `json:"email" validate:"required"`
	Password        string     `json:"password" validate:"required"`
	RememberToken   string     `json:"remember_token"`
	EmailVerifiedAt string     `json:"email_verified_at"`
	UserID          uint       `json:"user_id,omitempty" validate:"required"`
	CreatedAt       *time.Time `json:"created_at,omitempty"`
	UpdatedAt       *time.Time `json:"updated_at,omitempty"`
}
