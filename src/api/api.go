package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api")
	api.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello world")
	})
}
