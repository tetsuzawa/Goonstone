// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	"github.com/tetsuzawa/Goonstone/containers/api/cmd/server/controller"
	"github.com/tetsuzawa/Goonstone/containers/api/internal/core"
	"github.com/tetsuzawa/Goonstone/containers/api/pkg/awsx"
)

// Injectors from wire.go:

func InitializeControllers(db *gorm.DB, dbSessions redis.Conn, strg *awsx.Connection) *controller.Controllers {
	repository := core.NewGateway(db, dbSessions, strg)
	provider := core.NewProvider(repository)
	controllerController := controller.NewController(provider)
	controllers := controller.NewControllers(controllerController)
	return controllers
}
