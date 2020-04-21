package core

import (
	"context"
	"fmt"
	"github.com/tetsuzawa/Goonstone/containers/api/pkg/cerrors"
	"go.uber.org/multierr"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"sync"
)

// MockGateway - MockDBのアダプターの構造体
type MockGateway struct {
	db         *MockDB
	dbSessions map[string]uint //TODO session管理用にKVSの導入 (#18)
	storage    string
}

// NewMockGateway - MockDBのアダプターの構造体のコンストラクタ
func NewMockGateway(db *MockDB) Repository {
	strg, err := ioutil.TempDir("", "s3")
	if err != nil {
		log.Fatalln(err)
	}
	return &MockGateway{db, make(map[string]uint), strg}
}

// MockUsersTable - ユーザーテーブル
type MockUsersTable struct {
	data  map[uint]User
	index uint
}

func newMockUsersTable() *MockUsersTable {
	return &MockUsersTable{data: make(map[uint]User), index: 0}
}

// MockPhotosTable - 写真テーブル
type MockPhotosTable struct {
	data  map[string]Photo
	index string
}

func newMockPhotosTable() *MockPhotosTable {
	return &MockPhotosTable{data: make(map[string]Photo), index: ""}
}

// MockDB - テスト・開発用のDB
type MockDB struct {
	mu     sync.RWMutex
	users  *MockUsersTable
	photos *MockPhotosTable
}

func (mdb *MockDB) Close() {
}

// NewMockDB - テスト・開発用のDBのコンストラクタ
func NewMockDB() *MockDB {
	return &MockDB{users: newMockUsersTable(), photos: newMockPhotosTable()}
}

// MockDBSessions - テスト・開発用のセッション管理DB
type MockDBSessions struct {
	mu   sync.RWMutex
	data map[string]uint
}

// NewMockDBSessions - テスト・開発用のセッション管理DBのコンストラクタ
func NewMockDBSessions() *MockDBSessions {
	return &MockDBSessions{data: make(map[string]uint)}
}

// CreateUser - Userを登録
func (r *MockGateway) CreateUser(ctx context.Context, user User) (User, error) {
	r.db.mu.Lock()
	defer r.db.mu.Unlock()
	r.db.users.index++

	user.ID = r.db.users.index
	r.db.users.data[user.ID] = user

	return user, nil
}

// ReadUserByID - 指定したIDのユーザーを取得
func (r *MockGateway) ReadUserByID(ctx context.Context, id uint) (User, error) {
	user, ok := r.db.users.data[id]
	if !ok {
		return User{}, cerrors.ErrNotFound
	}
	return user, nil
}

// ReadUserByEmail - 指定したEmailのユーザーを取得
func (r *MockGateway) ReadUserByEmail(ctx context.Context, email string) (User, error) {
	// TODO: slow O(N)
	for _, u := range r.db.users.data {
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
func (r *MockGateway) CreatePhoto(ctx context.Context, user User, fileName string, file multipart.File, photo Photo) error {

	f, err := os.Create(filepath.Join(r.storage, fileName))
	if err != nil {
		err = multierr.Combine(err, cerrors.ErrInternal)
		return fmt.Errorf("os.Create: %w", err)
	}
	_, err = io.Copy(f, file)
	if err != nil {
		err = multierr.Combine(err, cerrors.ErrInternal)
		return fmt.Errorf("io.Copy: %w", err)
	}
	r.db.mu.Lock()
	defer r.db.mu.Unlock()
	r.db.photos.index = filepath.Base(fileName[:len(fileName)-len(filepath.Ext(fileName))])
	photo.ID = r.db.photos.index
	r.db.photos.data[photo.ID] = photo
	return nil
}
