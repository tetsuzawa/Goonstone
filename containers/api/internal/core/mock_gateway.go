package core

import (
	"context"
	"sync"

	"github.com/tetsuzawa/Goonstone/containers/api/pkg/cerrors"
)

// MockGateway - MockDBのアダプターの構造体
type MockGateway struct {
	db *MockDB
}

// MockDB - テスト・開発用のDB
type MockDB struct {
	mu    sync.RWMutex
	data  map[uint]User
	index uint
}

// NewMockDB - テスト・開発用のDBのコンストラクタ
func NewMockDB() *MockDB {
	return &MockDB{data: make(map[uint]User)}
}

// NewMockGateway - MockDBのアダプターの構造体のコンストラクタ
func NewMockGateway(db *MockDB) Repository {
	return &MockGateway{db}
}

// CreateUser - Userを登録
func (r *MockGateway) CreateUser(ctx context.Context, user User) (User, error) {
	r.db.mu.Lock()
	defer r.db.mu.Unlock()
	r.db.index++

	user.ID = r.db.index
	r.db.data[user.ID] = user

	return user, nil
}

// ReadUserByID - 指定したIDのユーザーを取得
func (r *MockGateway) ReadUserByID(ctx context.Context, id uint) (User, error) {
	user, ok := r.db.data[id]
	if !ok {
		return User{}, cerrors.ErrNotFound
	}
	return user, nil
}

// ReadUserByEmail - 指定したEmailのユーザーを取得
func (r *MockGateway) ReadUserByEmail(ctx context.Context, email string) (User, error) {
	// TODO: slow O(N)
	for _, u := range r.db.data {
		if u.Email == email {
			return u, nil
		}
	}
	return User{}, cerrors.ErrNotFound
}
