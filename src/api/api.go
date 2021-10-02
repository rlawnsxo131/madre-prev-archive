package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rlawnsxo131/madre-server/src/user"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api")
	api.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello world")
	})
	user.RegisterRoutes(api)
}
