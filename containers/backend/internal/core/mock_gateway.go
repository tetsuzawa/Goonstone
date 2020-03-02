package core

import (
	"sync"
)

// MockGateway TODO
type MockGateway struct {
	db *MockDB
}

// MockDB TODO
type MockDB struct {
	mu   sync.RWMutex
	data map[int64]User
}

// NewMockDB TODO
func NewMockDB() *MockDB {
	return &MockDB{data: make(map[int64]User)}
}

// NewMockGateway TODO
func NewMockGateway(db *MockDB) Repository {
	return &MockGateway{db}
}
