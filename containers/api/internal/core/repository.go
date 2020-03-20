package core

import "context"

// Repository - アプリケーションコアからDBへのアダプター
type Repository interface {
	CreateUser(ctx context.Context, user User) (User, error)
	ReadUserByID(ctx context.Context, id uint) (User, error)
	ReadUserByEmail(ctx context.Context, email string) (User, error)
}
