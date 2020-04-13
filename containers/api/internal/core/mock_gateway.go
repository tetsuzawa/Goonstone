package core

import (
	"context"
	"mime/multipart"
	"sync"

	"github.com/tetsuzawa/Goonstone/containers/api/pkg/cerrors"
)

// MockGateway - MockDBのアダプターの構造体
type MockGateway struct {
	db         *MockDB
	dbSessions map[string]uint //TODO session管理用にKVSの導入 (#18)
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

// MockDBSessions - テスト・開発用のセッション管理DB
type MockDBSessions struct {
	mu   sync.RWMutex
	data map[string]uint
}

// NewMockDB - テスト・開発用のセッション管理DBのコンストラクタ
func NewMockDBSessions() *MockDBSessions {
	return &MockDBSessions{data: make(map[string]uint)}
}

// NewMockGateway - MockDBのアダプターの構造体のコンストラクタ
func NewMockGateway(db *MockDB) Repository {
	return &MockGateway{db, make(map[string]uint)}
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

// CreateSessionBySessionIDUserID - セッションIDとユーザーIDからセッションを作成
func (r *MockGateway) CreateSessionBySessionIDUserID(ctx context.Context, sID string, id uint) error {
	//TODO
	r.dbSessions[sID] = id
	return nil
}

// ReadUserIDBySessionID - セッションIDからユーザーIDを取得
func (r *MockGateway) ReadUserIDBySessionID(ctx context.Context, sID string) (uint, error) {
	uID, ok := r.dbSessions[sID]
	if !ok {
		return 0, cerrors.ErrNotFound
	}
	return uID, nil
}

// CreatePhoto - 写真を保存する
func (r *MockGateway) CreatePhoto(ctx context.Context, fileName string, file multipart.File) error {
	//TODO
	return nil
}
