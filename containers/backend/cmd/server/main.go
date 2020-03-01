package main

import (
	"fmt"
	"github.com/tetsuzawa/goonstone/cmd/server/controllers"
	"github.com/tetsuzawa/goonstone/internal/core"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/tetsuzawa/goonstone/config"
)

func main() {
	e := newEcho()
	db := newDB()
	handler := newHandler(e, db)

	// Start server
	//e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%s", config.API.Host, config.API.Port)))
	log.Printf("Listening on port %s", config.API.Port)
	err := http.ListenAndServe(fmt.Sprintf("%s:%s", config.API.Host, config.API.Port), handler)
	if err != nil {
		log.Fatal(err)
	}

}

func newEcho() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())

	return e
}

func newDB() *gorm.DB {
	DBMS := config.DB.GormPrefix
	CONNECT := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		config.DB.User,
		config.DB.Password,
		config.DB.Host,
		config.DB.Port,
		config.DB.Database,
	)
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		log.Fatalln(err)
	}
	return db
}

func newHandler(e *echo.Echo, db *gorm.DB) http.Handler {
	gateway := core.NewGateway(db)
	provider := core.NewProvider(gateway)
	ctrl := controllers.NewController(provider)

	// Mux
	e.GET("/:message", ctrl.HandleMessage)

	return e
}
