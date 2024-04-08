package handlers

import (
	"github.com/labstack/echo/v4"

	"github.com/tommaso-merlini/go-templ-base-template/views/pages"
)

func HomeShow(c echo.Context) error {
	return render(c, pages.Index())
}
