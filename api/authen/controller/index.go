package controller

import (
	"net/http"
	"ticket/api"

	"github.com/labstack/echo/v4"
)

func Index(a *api.API) *echo.Echo {
	g := a.App

	g.POST("/sign-in", func(c echo.Context) error {
		return c.JSON(200, map[string]interface{}{
			"error":   false,
			"message": "sign-in",
		})
	})

	g.POST("/sign-up", func(c echo.Context) error {
		var body struct {
			Email    string `json:"email" validate:"required,email"`
			Password string `json:"password" validate:"required,min=8,max=32"`
		}

		err := c.Bind(&body)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		err = c.Validate(&body)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		return c.JSON(200, map[string]interface{}{
			"error":   false,
			"message": "sign-up",
		})
	})

	g.POST("/refresh-token", func(c echo.Context) error {
		return c.JSON(200, map[string]interface{}{
			"error":   false,
			"message": "refresh-token",
		})
	})

	return g
}