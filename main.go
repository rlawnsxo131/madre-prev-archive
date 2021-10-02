package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rlawnsxo131/madre-server/src/api"
)

func main() {
	e := echo.New()
	e.Debug = true
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	api.RegisterRoutes(e)
	e.Logger.Fatal(e.Start(":3001"))
}
