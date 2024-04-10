package views

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func Render(c echo.Context, component templ.Component) error {
	c.Response().Header().Set(echo.HeaderContentType, "text/html")
	return component.Render(c.Request().Context(), c.Response().Writer)
}

func HxRedirect(c echo.Context, to string) {
	c.Response().Header().Set("HX-Redirect", to)
}
