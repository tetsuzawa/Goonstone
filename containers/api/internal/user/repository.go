package user

import "context"

// Repository TODO
type Repository interface {
	CreateUser(ctx context.Context, user User) (User, error)
}
