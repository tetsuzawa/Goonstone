package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/tetsuzawa/Goonstone/config"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)

	fmt.Printf("\n\nServer is starting at: 80\n\n")
	// Start server
	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%s", config.API.Host, config.API.Port)))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
