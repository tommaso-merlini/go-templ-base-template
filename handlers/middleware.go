package handlers

import (
	"os"
	"strings"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"

	"github.com/tommaso-merlini/go-templ-base-template/auth"
	"github.com/tommaso-merlini/go-templ-base-template/db/queries"
	"github.com/tommaso-merlini/go-templ-base-template/shared"
)

const (
	sessionUserKey        = "user"
	sessionAccessTokenKey = "accessToken"
)

func WithAuthUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if strings.Contains(c.Request().URL.Path, "/public") ||
			strings.Contains(c.Request().URL.Path, "/images") {
			return next(c)
		}
		accessToken, err := getAccessToken(c)
		if err != nil {
			c.Set(sessionUserKey, shared.AuthUser{})
			return next(c)
		}
		_ = accessToken
		resp, err := auth.Client.Auth.User(c.Request().Context(), accessToken)
		if err != nil {
			c.Set(sessionUserKey, shared.AuthUser{})
			return next(c)
		}
		dbUser, err := queries.GetUserByAuthID(c.Request().Context(), resp.ID)
		if err != nil {
			c.Set(sessionUserKey, shared.AuthUser{})
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
		c.Set(sessionUserKey, user)
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
