package core

import (
	"context"
	"errors"

	"github.com/jinzhu/gorm"
	"go.uber.org/multierr"

	"github.com/tetsuzawa/Goonstone/containers/api/pkg/cerrors"
)

// Gateway - DBのアダプターの構造体
type Gateway struct {
	db *gorm.DB
}

// NewGateway - DBのアダプターの構造体のコンストラクタ
func NewGateway(db *gorm.DB) Repository {
	return &Gateway{db}
}

// CreateUser - ユーザーを登録
func (r *Gateway) CreateUser(ctx context.Context, user User) (User, error) {
	err := r.db.Create(&user).Error
	return user, err
}

// ReadUserByID - ユーザーを取得
func (r *Gateway) ReadUserByID(ctx context.Context, id uint) (User, error) {
	var user User
	err := r.db.First(&user, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = multierr.Combine(err, cerrors.ErrNotFound)
	}
	return user, err
}

// ReadUserByEmail - ユーザーを取得
func (r *Gateway) ReadUserByEmail(ctx context.Context, email string) (User, error) {
	var user User
	err := r.db.First(&user, email).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = multierr.Combine(err, cerrors.ErrNotFound)
	}
	return user, err
}
