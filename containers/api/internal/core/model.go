package core

import "time"

// Recipe - Userのモデル.
type User struct {
	ID                   uint       `json:"id,omitempty"`
	Name                 string     `json:"name,omitempty" validate:"required"`
	Email                string     `json:"email,omitempty" validate:"required"`
	Password             string     `json:"password,omitempty" validate:"required"`
	PasswordConfirmation string     `json:"password_confirmation,omitempty" validate:"required"`
	RememberToken        string     `json:"remember_token,omitempty"`
	EmailVerifiedAt      string     `json:"email_verified_at,omitempty"`
	CreatedAt            *time.Time `json:"created_at,omitempty"`
	UpdatedAt            *time.Time `json:"updated_at,omitempty"`
}
