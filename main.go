package main

import (
	"embed"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/tommaso-merlini/go-templ-base-template/handler"
)

//go:embed public
var FS embed.FS

func main() {
	e := echo.New()
	e.GET("/*", echo.WrapHandler(http.StripPrefix("/", http.FileServer(http.FS(FS)))))
	e.GET("/", handler.Make(handler.HomeShow))

	e.Logger.Fatal(e.Start(":3000"))
}
