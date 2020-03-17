package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
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
	Message string     `json:"message,omitempty"`
	User    *core.User `json:"user,omitempty"`
}

// HandleCreateRecipes - Ping用のルート.
// @Summary Ping用のルート.
// @Description Getでアクセスすると"OK"を返す
// @Accept json
// @Produce json
// @Success 200 {object} Response
// @Router /ping/ [get]
func (ctrl *Controller) HandlePing(c echo.Context) error {
	resp := Response{Message: "OK"}
	return c.JSON(http.StatusOK, resp)
}

// HandleRegisterUser - ユーザーを登録.
// @Summary ユーザーを登録.
// @Description title, making_tike, serves, ingredients, costからレシピを作成する
// @Accept json
// @Produce json
// @Param name query string true "Name"
// @Param email query string true "Email"
// @Param password query string true "password"
// @Param password_confirmation query string true "Password Confirmation"
// @Success 201 {object} Response
// @Failure 400 {object} Response
// @Failure 500 {object} Response
// @Router /recipes/ [post]
func (ctrl *Controller) HandleRegisterUser(c echo.Context) error {
	resp := Response{
		Message: "User registration failed",
		//User: &core.User{},
	}

	var in core.User
	if err := c.Bind(&in); err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, resp)
	}
	err := core.Validate.Struct(&in)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, resp)
	}
	if in.PasswordConfirmation != in.Password{
		err = fmt.Errorf("password does not match")
		log.Println(err)
		resp.Message = err.Error()
		return c.JSON(http.StatusBadRequest, resp)
	}

	ctx := c.Request().Context()
	user, err := ctrl.p.CreateUser(ctx, in)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, resp)
	}

	resp = Response{
		Message: "User successfully created!",
		User: &core.User{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
	}

	return c.JSON(http.StatusCreated, resp)
}
