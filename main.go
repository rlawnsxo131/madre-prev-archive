package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rlawnsxo131/madre-server/src/api"
	"github.com/rlawnsxo131/madre-server/src/graphql"
)

func main() {
	e := echo.New()
	e.Debug = true
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	api.RegisterRoutes(e)

	// graphql
	handler, err := graphql.NewHandler()
	if err != nil {
		e.Logger.Fatal(err)
	}
	e.GET("/graphql", echo.WrapHandler(handler))
	e.POST("/graphql", echo.WrapHandler(handler))

	e.Logger.Fatal(e.Start(":3001"))
}
