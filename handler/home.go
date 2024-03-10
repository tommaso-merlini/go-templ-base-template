package handler

import (
	"github.com/labstack/echo/v4"

	"github.com/tommaso-merlini/go-templ-base-template/view/home"
)

func HomeShow(c echo.Context) error {
	return render(c, home.Home())
}
