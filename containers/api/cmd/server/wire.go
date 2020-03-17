//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"

	"github.com/tetsuzawa/Goonstone/containers/api/cmd/server/controller"
	"github.com/tetsuzawa/Goonstone/containers/api/internal/core"
)

// InitializeControllers - 依存管理. wireでDIする.
func InitializeControllers(db *gorm.DB) *controller.Controllers {
	wire.Build(
		core.NewGateway,
		core.NewProvider,
		controller.NewController,
		controller.NewControllers,
	)
	return &controller.Controllers{}
}