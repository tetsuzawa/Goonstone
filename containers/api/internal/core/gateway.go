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
	db         *gorm.DB
	dbSessions map[string]uint //TODO session管理用にKVSの導入 (#18)
}

// NewGateway - DBのアダプターの構造体のコンストラクタ
func NewGateway(db *gorm.DB) Repository {
	return &Gateway{db, map[string]uint{}}
}

// CreateUser - ユーザーを登録
func (r *Gateway) CreateUser(ctx context.Context, user User) (User, error) {
	//TODO uniqueのエラーハンドリング (#17)
	if !r.db.NewRecord(user) {
		return User{}, cerrors.ErrAlreadyExists
	}
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

// CreateSessionBySessionIDUserID - セッションIDとユーザーIDからセッションを作成
func (r *Gateway) CreateSessionBySessionIDUserID(ctx context.Context, sID string, id uint) error {
	//TODO
	r.dbSessions[sID] = id
	return nil
}

// ReadUserIDBySessionID - セッションIDからユーザーIDを取得
func (r *Gateway) ReadUserIDBySessionID(ctx context.Context, sID string) (uint, error) {
	uID, ok := r.dbSessions[sID]
	if !ok {
		err := errors.New("failed to read user id from dbSessions")
		err = multierr.Combine(err, cerrors.ErrNotFound)
		return 0, err
	}
	return uID, nil
}
