package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"

	"github.com/tetsuzawa/Goonstone/containers/api/internal/core"
)

// Controller - リクエストを処理しアプリケーションコアに渡す
type Controller struct {
	p *core.Provider
}

// NewController - Controllerのコンストラクタ
func NewController(p *core.Provider) *Controller {
	return &Controller{p}
}

// Response - Controllerのレスポンスを定義した構造体
type Response struct {
	Message string `json:"message,omitempty"`
}

// HandlePing TODO
func (ctrl *Controller) HandlePing(c echo.Context) error {
	resp := Response{Message: "OK"}
	return c.JSON(http.StatusOK, resp)
}
