package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/tetsuzawa/Goonstone/containers/api/pkg/cerrors"
	"go.uber.org/multierr"
	"net/http"
	"time"
)

const (
	sessionExpiresAt = 24 * 7 * time.Hour
	sessionCookieName = "session"
)

func WriteSessionCookie(c echo.Context, sID string) error {
	cookie := &http.Cookie{
		Name:    sessionCookieName,
		Value:   sID,
		Expires: time.Now().Add(sessionExpiresAt),
		Path:    "/",
	}
	c.SetCookie(cookie)
	return nil
}

func ReadSessionCookie(c echo.Context) (string, error) {
	cookie, err := c.Cookie(sessionCookieName)
	if err == echo.ErrCookieNotFound {
		return "", multierr.Combine(err, cerrors.ErrNotFound)
	} else if err == http.ErrNoCookie {
		return "", multierr.Combine(err, cerrors.ErrNotFound)
	} else if err != nil {
		return "", multierr.Combine(err, cerrors.ErrInternal)
	}
	cookie.Expires = time.Now().Add(sessionExpiresAt)
	c.SetCookie(cookie)
	return cookie.Value, nil
}
