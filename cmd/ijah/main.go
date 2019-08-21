package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/taufanmahaputra/ijah/pkg/lib/config"
	"github.com/taufanmahaputra/ijah/pkg/service"
)

func main() {
	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	server := service.NewHTTPServer()
	server.RegisterHandler(e)

	cfg := config.GetConfig()

	e.Logger.Fatal(e.Start(cfg.App.Port))
}
