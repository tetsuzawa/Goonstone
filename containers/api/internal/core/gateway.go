package core

import (
	"context"

	"github.com/jinzhu/gorm"
)

// Gateway - DBのアダプターの構造体
type Gateway struct {
	db *gorm.DB
}

// NewGateway - DBのアダプターの構造体のコンストラクタ
func NewGateway(db *gorm.DB) Repository {
	return &Gateway{db}
}

// CreateUser - Userを登録
func (r *Gateway) CreateUser(ctx context.Context, user User) (User, error) {
	err := r.db.Create(&user).Error
	return user, err
}

// ReadUser - 指定したIDのユーザーを取得
func (r *Gateway) ReadUser(ctx context.Context, user User) (User, error) {
	err := r.db.First(&user, user.Email).Error
	return user, err
}
