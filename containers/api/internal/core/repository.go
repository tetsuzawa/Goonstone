package core

import "context"

// Repository - アプリケーションコアからDBへのアダプター
type Repository interface {
	CreateUser(ctx context.Context, user User) (User, error)
	ReadUser(ctx context.Context, user User) (User, error)
}
