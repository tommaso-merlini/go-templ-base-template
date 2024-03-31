package handler

import (
	"os"
	"strings"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"

	"github.com/tommaso-merlini/go-templ-base-template/db"
	sb "github.com/tommaso-merlini/go-templ-base-template/pkg/supaauth"
	"github.com/tommaso-merlini/go-templ-base-template/shared"
)

const (
	sessionUserKey        = "user"
	sessionAccessTokenKey = "accessToken"
)

func WithAuthUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if strings.Contains(c.Request().URL.Path, "/public") {
			return next(c)
		}
		accessToken, err := getAccessToken(c)
		if err != nil {
			c.Set("user", shared.AuthUser{})
			return next(c)
		}
		_ = accessToken
		resp, err := sb.Client.Auth.User(c.Request().Context(), accessToken)
		if err != nil {
			c.Set("user", shared.AuthUser{})
			return next(c)
		}
		dbUser, err := db.GetUserByAuthID(c.Request().Context(), resp.ID)
		if err != nil {
			c.Set("user", shared.AuthUser{})
			return next(c)
		}
		user := shared.AuthUser{
			ID:         dbUser.ID.String(),
			AuthID:     resp.ID,
			Email:      resp.Email,
			Active:     dbUser.Active,
			CreatedAt:  dbUser.CreatedAt,
			IsLoggedIn: true,
		}
		c.Set("user", user)
		return next(c)
	}
}

func getAccessToken(c echo.Context) (string, error) {
	store := sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))
	session, err := store.Get(c.Request(), sessionUserKey)
	if err != nil {
		return "", err
	}
	accessToken := session.Values[sessionAccessTokenKey]
	if accessToken == nil {
		return "", nil
	}
	return accessToken.(string), nil
}
