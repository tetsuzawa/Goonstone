package controller

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.uber.org/multierr"

	"github.com/tetsuzawa/Goonstone/containers/api/pkg/cerrors"
)

const (
	sessionExpiration = 24 * 7 * time.Hour
	sessionCookieName = "session"
)

func WriteSessionIDToCookie(c echo.Context, sID string) {
	cookie := &http.Cookie{
		Name:    sessionCookieName,
		Value:   sID,
		Expires: time.Now().Add(sessionExpiration),
		Path:    "/",
	}
	c.SetCookie(cookie)
}

func ReadSessionIDFromCookie(c echo.Context) (string, error) {
	cookie, err := c.Cookie(sessionCookieName)
	if err == echo.ErrCookieNotFound {
		return "", multierr.Combine(err, cerrors.ErrNotFound)
	} else if err == http.ErrNoCookie {
		return "", multierr.Combine(err, cerrors.ErrNotFound)
	} else if err != nil {
		return "", multierr.Combine(err, cerrors.ErrInternal)
	}
	return cookie.Value, nil
}

func DeleteSessionIDFromCookie(c echo.Context) {
	cookie := &http.Cookie{
		Name:    sessionCookieName,
		Value:   "",
		Expires: time.Unix(0, 0),
		Path:    "/",
	}
	c.SetCookie(cookie)
}
