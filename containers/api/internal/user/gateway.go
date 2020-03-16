package user

import (
	"context"
	"github.com/jinzhu/gorm"
)

// Gateway TODO
type Gateway struct {
	db *gorm.DB
}

// NewGateway TODO
func NewGateway(db *gorm.DB) Repository {
	return &Gateway{db}
}

// CreateRecipe - Userを登録
func (r *Gateway) CreateUser(ctx context.Context, user User) (User, error) {
	err := r.db.Create(&user).Error
	return user, err
}
