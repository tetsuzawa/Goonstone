package core

import "time"

// User - Userのモデル.
type User struct {
	ID                   uint       `json:"id,omitempty"`
	Name                 string     `json:"name,omitempty" validate:"required"`
	Email                string     `json:"email,omitempty" validate:"required"`
	Password             string     `json:"password,omitempty" validate:"required"`
	PasswordConfirmation string     `json:"password_confirmation,omitempty" validate:"required" gorm:"-"`
	RememberToken        string     `json:"remember_token,omitempty"`
	EmailVerifiedAt      *time.Time `json:"email_verified_at,omitempty"`
	CreatedAt            *time.Time `json:"created_at,omitempty"`
	UpdatedAt            *time.Time `json:"updated_at,omitempty"`
}

// Photo - Photoのモデル.
type Photo struct {
	ID        string     `json:"id,omitempty"`
	UserID    uint       `json:"user_id,omitempty"`
	FileName  string     `json:"file_name,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// Comment - Commentのモデル.
type Comment struct {
	ID        uint       `json:"id,omitempty"`
	PhotoID   uint       `json:"photo_id,omitempty"`
	UserID    uint       `json:"user_id,omitempty"`
	Content   string     `json:"content,omitempty" validate:"required"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// Like - Likeのモデル.
type Like struct {
	ID        uint       `json:"id,omitempty"`
	PhotoID   uint       `json:"photo_id,omitempty"`
	UserID    uint       `json:"user_id,omitempty"`
	Content   string     `json:"content,omitempty" validate:"required"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}
