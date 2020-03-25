package core

import (
	"context"
	"errors"
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	"go.uber.org/multierr"

	"github.com/tetsuzawa/Goonstone/containers/api/pkg/cerrors"
)

// Gateway - DBのアダプターの構造体
type Gateway struct {
	db         *gorm.DB
	dbSessions redis.Conn //TODO session管理用にKVSの導入 (#18)
}

// NewGateway - DBのアダプターの構造体のコンストラクタ
func NewGateway(db *gorm.DB, conn redis.Conn) Repository {
	return &Gateway{db, conn}
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
	err := r.db.Where("email = ?", email).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = multierr.Combine(err, cerrors.ErrNotFound)
	} else if err != nil {
		err = multierr.Combine(err, cerrors.ErrInternal)
	}
	return user, err
}

// CreateSessionBySessionIDUserID - セッションIDとユーザーIDからセッションを作成
func (r *Gateway) CreateSessionBySessionIDUserID(ctx context.Context, sID string, id uint) error {
	reply, err := r.dbSessions.Do("SET", sID, id)
	if reply != "OK" || err != nil {
		return multierr.Combine(err, cerrors.ErrInternal)
	}
	return nil
}

// ReadUserIDBySessionID - セッションIDからユーザーIDを取得
func (r *Gateway) ReadUserIDBySessionID(ctx context.Context, sID string) (uint, error) {
	uID, err := redis.Uint64(r.dbSessions.Do("GET", sID))
	if errors.Is(err, redis.ErrNil) {
		return 0, multierr.Combine(err, cerrors.ErrNotFound)
	} else if err != nil {
		return 0, multierr.Combine(err, cerrors.ErrInternal)
	}
	return uint(uID), nil
}
