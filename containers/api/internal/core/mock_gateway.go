package core

import (
	"context"
	"sync"
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
