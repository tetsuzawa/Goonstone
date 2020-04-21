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
	"github.com/tetsuzawa/Goonstone/containers/api/pkg/awsx"
	"github.com/tetsuzawa/Goonstone/containers/api/pkg/mysql"
	"github.com/tetsuzawa/Goonstone/containers/api/pkg/redisx"
	"github.com/tetsuzawa/Goonstone/containers/api/pkg/webcfg"
)

const applicationURL = "goonstone.tetsuzawa.com"

// @title Goonstone - Picture sharing web-app written in Go
// @version 1.0
// @description This is a recipes API server.
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host goonstone.tetsuzawa.com:80
// @BasePath /
func main() {
	// default
	var apiCfg webcfg.APIConfig
	err := webcfg.ReadAPIEnv(&apiCfg)
	if err != nil {
		log.Fatalf("%+v", err)
	}

	db := newDB()
	defer db.Close()
	dbSessions := newDBSessions()
	defer dbSessions.Close()
	strg := newStorage()
	ctrls := InitializeControllers(db, dbSessions, strg)
	handler := newHandler(ctrls)

	log.Printf("Listening on %s:%s", apiCfg.Host, apiCfg.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", apiCfg.Host, apiCfg.Port), handler))
}

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func newDB() *gorm.DB {
	// Mysql
	var cfg mysql.Config
	err := mysql.ReadEnv(&cfg)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Connecting to MYSQL ...")
	db, err := mysql.Connect(cfg)
	if err != nil {
		log.Fatalln(err)
	}
	return db
}

func newDBSessions() redis.Conn {
	// Redis
	var cfg redisx.Config
	err := redisx.ReadEnv(&cfg)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Connecting to REDIS ...")
	dbSessions, err := redisx.Connect(cfg)
	if err != nil {
		log.Fatalln(err)
	}
	return dbSessions
}

func newStorage() *awsx.Connection {
	var cfg awsx.Config
	err := awsx.ReadEnv(&cfg)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Connecting to S3 ...")
	strg, err := awsx.Connect(cfg)
	return strg
}

func newHandler(ctrls *controller.Controllers) http.Handler {
	var frontendCfg webcfg.FRONTENDConfig
	err := webcfg.ReadFRONTENDEnv(&frontendCfg)
	if err != nil {
		log.Fatalf("%+v", err)
	}

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	//e.Use(middleware.CSRF())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: true,
		AllowOrigins: []string{
			fmt.Sprintf("http://%s", applicationURL),
			fmt.Sprintf("http://%s", frontendCfg.Host),
			fmt.Sprintf("http://%s:%s", frontendCfg.Host, frontendCfg.Port),
			"http://127.0.0.1",
		},
	}))

	api := e.Group("/api")
	api.GET("/ping", ctrls.Ctrl.HandlePing)
	api.POST("/register", ctrls.Ctrl.HandleRegisterUser)
	api.POST("/login", ctrls.Ctrl.HandleLoginUser)
	api.POST("/logout", ctrls.Ctrl.HandleLogoutUser)
	api.GET("/user", ctrls.Ctrl.HandleReadUserDetails)
	api.POST("/photos", ctrls.Ctrl.HandleStorePhoto)
	// swagger
	api.GET("/swagger/*", echoSwagger.WrapHandler)
	return e
}
