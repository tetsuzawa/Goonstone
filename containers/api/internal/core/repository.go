package core

import (
	"context"
	"mime/multipart"
)

// Repository - アプリケーションコアからDBへのアダプター
type Repository interface {
	CreateUser(ctx context.Context, user User) (User, error)
	ReadUserByID(ctx context.Context, id uint) (User, error)
	ReadUserByEmail(ctx context.Context, email string) (User, error)
	CreateSessionBySessionIDUserID(ctx context.Context, sID string, id uint) error
	ReadUserIDBySessionID(ctx context.Context, sID string) (uint, error)
	CreatePhoto(ctx context.Context, user User, fileName string, file multipart.File, photo Photo) error
}
