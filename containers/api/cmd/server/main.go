package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/swaggo/echo-swagger"

	"github.com/tetsuzawa/Goonstone/containers/api/cmd/server/controller"
	"github.com/tetsuzawa/Goonstone/containers/api/pkg/env"
	"github.com/tetsuzawa/Goonstone/containers/api/pkg/mysql"
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
		log.Println(err)
		apiCfg.Host = "127.0.0.1"
		apiCfg.Port = "8080"
	}
	db := newDB()
	ctrls := InitializeControllers(db)
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

func newHandler(e *echo.Echo, ctrls *controller.Controllers) http.Handler {
	api := e.Group("/api")
	api.GET("/ping/", ctrls.Ctrl.HandlePing)
	api.POST("/register/", ctrls.Ctrl.HandleRegisterUser)
	// swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	return e
}
