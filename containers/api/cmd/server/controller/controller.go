package controller

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/labstack/echo/v4"

	"github.com/tetsuzawa/Goonstone/containers/api/internal/core"
	"github.com/tetsuzawa/Goonstone/containers/api/pkg/cerrors"
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
	Message string        `json:"message,omitempty"`
	User    *core.User    `json:"user,omitempty"`
	Photo   *core.Photo   `json:"photo,omitempty"`
	Comment *core.Comment `json:"comment,omitempty"`
	Like    *core.Like    `json:"like,omitempty"`
}

// HandleCreateRecipes - Ping用のルート.
// @Summary Ping用のルート.
// @Description Getでアクセスすると"OK"を返す
// @Accept json
// @Produce json
// @Success 200 {object} Response
// @Router /ping [get]
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
// @Router /register [post]
func (ctrl *Controller) HandleRegisterUser(c echo.Context) error {
	resp := Response{
		Message: "User registration failed",
	}
	sID, err := ReadSessionIDFromCookie(c)
	if !errors.Is(err, cerrors.ErrNotFound) && err != nil {
		log.Printf("%+v", err)
		return c.JSON(http.StatusInternalServerError, resp)
	}
	ctx := c.Request().Context()
	alreadyLoggedIn, err := ctrl.p.AlreadyLoggedIn(ctx, sID)
	if err != nil {
		log.Printf("%+v", err)
		return c.JSON(http.StatusInternalServerError, resp)
	}
	if alreadyLoggedIn {
		WriteSessionIDToCookie(c, sID)
		resp.Message = "User already logged in"
		return c.JSON(http.StatusSeeOther, resp)
	}

	var in core.User
	if err := c.Bind(&in); err != nil {
		log.Printf("%+v", err)
		return c.JSON(http.StatusBadRequest, resp)
	}
	err = core.Validate.Struct(&in)
	if err != nil {
		log.Printf("%+v", err)
		return c.JSON(http.StatusBadRequest, resp)
	}
	if in.PasswordConfirmation != in.Password {
		err = fmt.Errorf("password does not match")
		log.Printf("%+v", err)
		resp.Message = err.Error()
		return c.JSON(http.StatusBadRequest, resp)
	}
	user, err := ctrl.p.CreateUser(ctx, in)
	if err != nil {
		//TODO: ErrInternalを使うか検討
		log.Printf("%+v", err)
		return c.JSON(http.StatusInternalServerError, resp)
	}
	sID, err = ctrl.p.CreateSession(ctx, user.ID)
	if err != nil {
		log.Printf("%+v", err)
		return c.JSON(http.StatusInternalServerError, resp)
	}
	WriteSessionIDToCookie(c, sID)

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

// HandleLoginUser - ログイン処理.
// @Summary ログイン処理.
// @Description email, passwordからユーザーをログイン処理する
// @Accept json
// @Produce json
// @Param email query string true "Email"
// @Param password query string true "password"
// @Success 200 {object} Response
// @Failure 303 {object} Response
// @Failure 400 {object} Response
// @Failure 401 {object} Response
// @Failure 404 {object} Response
// @Failure 500 {object} Response
// @Router /login [post]
func (ctrl *Controller) HandleLoginUser(c echo.Context) error {
	sID, err := ReadSessionIDFromCookie(c)
	if !errors.Is(err, cerrors.ErrNotFound) && err != nil {
		log.Printf("%+v", err)
		return c.JSON(http.StatusInternalServerError, Response{Message: "Login failed"})
	}
	ctx := c.Request().Context()
	alreadyLoggedIn, err := ctrl.p.AlreadyLoggedIn(ctx, sID)
	if err != nil {
		log.Printf("%+v", err)
		return c.JSON(http.StatusInternalServerError, Response{Message: "Login failed"})
	}
	if alreadyLoggedIn {
		WriteSessionIDToCookie(c, sID)
		return c.JSON(http.StatusSeeOther, Response{Message: "User already logged in"})
	}

	var in core.User
	if err := c.Bind(&in); err != nil {
		log.Printf("%+v", err)
		return c.JSON(http.StatusBadRequest, Response{Message: "Request is not valid"})
	}
	user, err := ctrl.p.LoginUser(ctx, in)
	if errors.Is(err, cerrors.ErrNotFound) {
		return c.JSON(http.StatusNotFound, Response{Message: "User does not exist"})
	} else if errors.Is(err, cerrors.ErrUnauthenticated) {
		return c.JSON(http.StatusUnauthorized, Response{Message: "Password is invalid"})
	} else if err != nil {
		log.Printf("%+v", err)
		return c.JSON(http.StatusInternalServerError, Response{Message: "Internal server error"})
	}
	sID, err = ctrl.p.CreateSession(ctx, user.ID)
	if err != nil {
		log.Printf("%+v", err)
		return c.JSON(http.StatusInternalServerError, Response{Message: "Internal server error"})
	}
	WriteSessionIDToCookie(c, sID)

	resp := Response{
		Message: "Successfully logged in!",
		User: &core.User{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
	}
	return c.JSON(http.StatusOK, resp)
}

// HandleLogoutUser - ログアウト処理.
// @Summary ログアウト処理.
// @Description CookieのセッションIDをもとにユーザーをログアウト処理する
// @Accept json
// @Produce json
// @Success 200 {object} Response
// @Failure 303 {object} Response
// @Failure 400 {object} Response
// @Failure 500 {object} Response
// @Router /logout [post]
func (ctrl *Controller) HandleLogoutUser(c echo.Context) error {
	sID, err := ReadSessionIDFromCookie(c)
	if !errors.Is(err, cerrors.ErrNotFound) && err != nil {
		log.Printf("%+v", err)
		return c.JSON(http.StatusInternalServerError, Response{Message: "Logout failed"})
	}
	ctx := c.Request().Context()
	alreadyLoggedIn, err := ctrl.p.AlreadyLoggedIn(ctx, sID)
	if err != nil {
		log.Printf("%+v", err)
		return c.JSON(http.StatusInternalServerError, Response{Message: "Logout failed"})
	}
	if !alreadyLoggedIn {
		return c.JSON(http.StatusSeeOther, Response{Message: "User has not logged in"})
	}
	DeleteSessionIDFromCookie(c)
	return c.JSON(http.StatusOK, Response{Message: "User Successfully logged out!"})
}

// HandleReadUserDetails - ユーザー詳細取得処理.
// @Summary ユーザー詳細取得処理.
// @Description CookieのセッションIDをもとにユーザー詳細を取得する
// @Accept json
// @Produce json
// @Success 200 {object} Response
// @Failure 303 {object} Response
// @Failure 500 {object} Response
// @Router /user [get]
func (ctrl *Controller) HandleReadUserDetails(c echo.Context) error {
	sID, err := ReadSessionIDFromCookie(c)
	if !errors.Is(err, cerrors.ErrNotFound) && err != nil {
		log.Printf("%+v", err)
		return c.JSON(http.StatusInternalServerError, Response{Message: "Read user details failed"})
	}
	ctx := c.Request().Context()
	alreadyLoggedIn, err := ctrl.p.AlreadyLoggedIn(ctx, sID)
	if err != nil {
		log.Printf("%+v", err)
		return c.JSON(http.StatusInternalServerError, Response{Message: "Login failed"})
	}
	if !alreadyLoggedIn {
		return c.JSON(http.StatusSeeOther, Response{Message: "User has not logged in"})
	}
	user, err := ctrl.p.ReadUserDetails(ctx, sID)
	if errors.Is(err, cerrors.ErrInternal) {
		log.Printf("%+v", err)
		return c.JSON(http.StatusInternalServerError, Response{Message: "Internal server error"})
	}
	return c.JSON(http.StatusOK, Response{
		Message: "User details successfully read!",
		User: &core.User{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
	})
}

// HandleStorePhoto - 写真をアップロード.
// @Summary 写真をアップロード.
// @Description 写真を登録する
// @Accept json
// @Produce json
// @Success 201 {object} Response
// @Failure 303 {object} Response
// @Failure 400 {object} Response
// @Failure 500 {object} Response
// @Router /photos [post]
func (ctrl *Controller) HandleStorePhoto(c echo.Context) error {
	sID, err := ReadSessionIDFromCookie(c)
	if !errors.Is(err, cerrors.ErrNotFound) && err != nil {
		log.Printf("%+v", err)
		return c.JSON(http.StatusInternalServerError, Response{Message: "Internal server error"})
	}
	ctx := c.Request().Context()
	alreadyLoggedIn, err := ctrl.p.AlreadyLoggedIn(ctx, sID)
	if err != nil {
		log.Printf("%+v", err)
		return c.JSON(http.StatusInternalServerError, Response{Message: "Internal server error"})
	}
	if !alreadyLoggedIn {
		return c.JSON(http.StatusSeeOther, Response{Message: "User has not logged in"})
	}
	photoFile, err := c.FormFile("photo")
	if errors.Is(err, http.ErrMissingFile) {
		log.Printf("%+v", err)
		return c.JSON(http.StatusUnprocessableEntity, Response{Message: "File not included"})
	} else if err != nil {
		log.Printf("%+v", err)
		return c.JSON(http.StatusInternalServerError, Response{Message: "Internal server error"})
	}
	ext := filepath.Ext(photoFile.Filename)
	if ext != ".png" &&
		ext != ".PNG" &&
		ext != ".jpg" &&
		ext != ".JPG" &&
		ext != ".gif" &&
		ext != ".GIF" {
		log.Printf("%+v", err)
		return c.JSON(http.StatusBadRequest, Response{Message: "Invalid file type"})
	}
	photo, err := ctrl.p.StorePhoto(ctx, sID, photoFile)
	if errors.Is(err, cerrors.ErrInternal) {
		log.Printf("%+v", err)
		return c.JSON(http.StatusInternalServerError, Response{Message: "Internal server error"})
	}

	return c.JSON(http.StatusCreated, Response{
		Message: "Photo successfully stored!",
		Photo:   &photo,
	})
}
