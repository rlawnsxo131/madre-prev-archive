package user

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func getUserHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	if id == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "not found user")
	}

	user := FindById(id)

	return c.JSON(http.StatusOK, user)
}

func RegisterRoutes(e *echo.Group) {
	user := e.Group("/user")
	user.GET("/:id", getUserHandler)
}
