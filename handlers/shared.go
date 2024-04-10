package handlers

import (
	"log/slog"

	"github.com/labstack/echo/v4"
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
