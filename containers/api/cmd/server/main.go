package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/swaggo/echo-swagger"

	"github.com/tetsuzawa/Goonstone/containers/api/cmd/server/controller"
	"github.com/tetsuzawa/Goonstone/containers/api/pkg/env"
	"github.com/tetsuzawa/Goonstone/containers/api/pkg/mysql"
	"github.com/tetsuzawa/Goonstone/containers/api/pkg/redisx"
)

// @title Goonstone - Picture sharing web-app written in Go
// @version 1.0
// @description This is a recipes API server.
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host goonstone.tetsuzawa.com:80
// @BasePath /api
func main() {
	e := createMux()
	apiCfg, err := env.ReadAPIEnv()
	if err != nil {
		log.Printf("%+v", err)
		apiCfg.Host = "127.0.0.1"
		apiCfg.Port = "8080"
	}
	db := newDB()
	defer db.Close()
	dbSessions := newDBSessions()
	defer dbSessions.Close()

	ctrls := InitializeControllers(db, dbSessions)
	handler := newHandler(e, ctrls)

	log.Printf("Listening on %s:%s", apiCfg.Host, apiCfg.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", apiCfg.Host, apiCfg.Port), handler))
}

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func createMux() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	return e
}

func newDB() *gorm.DB {
	// Mysql
	mysqlCfg, err := env.ReadMysqlEnv()
	if err != nil {
		log.Fatalln(err)
	}
	db, err := mysql.Connect(mysqlCfg)
	if err != nil {
		log.Fatalln(err)
	}
	return db
}

func newDBSessions() redis.Conn {
	// Redis
	redisCfg, err := env.ReadRedisEnv()
	if err != nil {
		log.Fatalln(err)
	}
	dbSessions, err := redisx.Connect(redisCfg)
	if err != nil {
		log.Fatalln(err)
	}
	return dbSessions
}

func newHandler(e *echo.Echo, ctrls *controller.Controllers) http.Handler {
	e.GET("/ping", ctrls.Ctrl.HandlePing)
	e.POST("/register", ctrls.Ctrl.HandleRegisterUser)
	e.POST("/login", ctrls.Ctrl.HandleLoginUser)
	e.POST("/logout", ctrls.Ctrl.HandleLogoutUser)
	//api := e.Group("/api")
	// swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	return e
}
