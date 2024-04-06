package handlers

import (
	"log/slog"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"

	"github.com/tommaso-merlini/go-templ-base-template/shared"
)

func Make(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := h(c)
		if err != nil {
			slog.Error("internal server error", "err", err, "path", c.Request().URL.Path)
		}
		return err
	}
}

func render(c echo.Context, component templ.Component) error {
	c.Response().Header().Set(echo.HeaderContentType, "text/html")
	return component.Render(c.Request().Context(), c.Response().Writer)
}

func getAuthUser(c echo.Context) (shared.AuthUser, bool) {
	u, ok := c.Get("user").(shared.AuthUser)
	if !ok {
		return shared.AuthUser{}, false
	}
	return u, u.IsLoggedIn
}

func hxRedirect(c echo.Context, to string) {
	c.Response().Header().Set("HX-Redirect", to)
}
