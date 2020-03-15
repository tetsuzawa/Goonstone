package core

import (
	"github.com/jinzhu/gorm"
)

// Gateway TODO
type Gateway struct {
	db *gorm.DB
}

// NewGateway TODO
func NewGateway(db *gorm.DB) Repository {
	return &Gateway{db}
}
