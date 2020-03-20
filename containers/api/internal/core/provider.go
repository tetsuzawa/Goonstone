package core

import (
	"context"
	"errors"
	"fmt"

	"go.uber.org/multierr"
	"golang.org/x/crypto/bcrypt"

	"github.com/tetsuzawa/Goonstone/containers/api/pkg/cerrors"
)

// Provider - アプリケーションコアの構造体
type Provider struct {
	r Repository
}

// NewProvider - アプリケーションコアの構造体のコンストラクタ
func NewProvider(r Repository) *Provider {
	return &Provider{r}
}

// hashPassword - パスワードハッシュを作る
func hashPassword(pw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), err
}

// verifyPassword - パスワードがハッシュにマッチするかどうかを調べる
func verifyPassword(hash, pw string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw))
}

// CreateUser - ユーザーを登録
func (p *Provider) CreateUser(ctx context.Context, user User) (User, error) {
	var err error
	user.Password, err = hashPassword(user.Password)
	if err != nil {
		err = multierr.Combine(cerrors.ErrInternal)
		err = fmt.Errorf("hashPassword: %w", err)
		return User{}, err
	}
	user, err = p.r.CreateUser(ctx, user)
	if err != nil {
		err = multierr.Combine(cerrors.ErrInternal)
		err = fmt.Errorf("CreateUser: %w", err)
		return User{}, err
	}
	return user, nil
}

// LoginUser - ユーザーのログイン処理
func (p *Provider) LoginUser(ctx context.Context, user User) (User, error) {
	reqPw := user.Password
	user, err := p.r.ReadUserByEmail(ctx, user.Email)
	if errors.Is(err, cerrors.ErrNotFound) {
		return User{}, err
	} else if err != nil {
		err = multierr.Combine(err, cerrors.ErrInternal)
		err = fmt.Errorf("ReadUserByEmail: %w", err)
		return User{}, err
	}
	err = verifyPassword(user.Password, reqPw)
	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		err = multierr.Combine(err, cerrors.ErrUnauthenticated)
		return User{}, err
	} else if err != nil {
		err = multierr.Combine(err, cerrors.ErrInternal)
		err = fmt.Errorf("verifyPassword: %w", err)
		return User{}, err
	}
	return user, nil
}
