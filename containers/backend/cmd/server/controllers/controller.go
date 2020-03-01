package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/tetsuzawa/goonstone/internal/core"
	"net/http"
)

// Controller TODO
type Controller struct {
	p *core.Provider
}

// NewController TODO
func NewController(p *core.Provider) *Controller {
	return &Controller{p}
}

// HandleMessage TODO
func (ctrl *Controller) HandleMessage(c echo.Context) error {
	msg := c.Param("message")
	if msg == "" {
		msg = "Hello"
	}
	return c.String(http.StatusOK, msg)
}
