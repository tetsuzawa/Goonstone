package core

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	"github.com/tetsuzawa/Goonstone/containers/api/pkg/awsx"
	"go.uber.org/multierr"
	"mime/multipart"

	"github.com/tetsuzawa/Goonstone/containers/api/pkg/cerrors"
)

// Gateway - DBのアダプターの構造体
type Gateway struct {
	db         *gorm.DB
	dbSessions redis.Conn //TODO session管理用にKVSの導入 (#18)
	storage    *awsx.Connection
}

// NewGateway - DBのアダプターの構造体のコンストラクタ
func NewGateway(db *gorm.DB, conn redis.Conn, strg *awsx.Connection) Repository {
	return &Gateway{db, conn, strg}
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

// CreatePhoto - 写真を保存する
func (r *Gateway) CreatePhoto(ctx context.Context, user User, fileName string, file multipart.File, photo Photo) error {
	uploader := s3manager.NewUploader(r.storage.Session)
	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(r.storage.Config.S3Bucket),
		Key:    aws.String(fileName),
		Body:   file,
	})
	if err != nil {
		return multierr.Combine(err, cerrors.ErrInternal)
	}
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&photo).Error; err != nil {
			err = fmt.Errorf("tx.Create: %w", err)
			err = multierr.Combine(cerrors.ErrInternal)
			_, err2 := r.storage.SVC.DeleteObject(&s3.DeleteObjectInput{
				Bucket: aws.String(r.storage.Config.S3Bucket),
				Key:    aws.String(fileName),
			})
			if err2 != nil {
				err2 = fmt.Errorf("s3.DeleteObject: %w", err)
				err = multierr.Combine(err, err2)
			}
			return err
		}
		return nil
	})
}
