package middlewareC

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v5"
)

func setTokenHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c *echo.Context) error {
		err := next(c)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, "Error")
		}

		token, ok := c.Get("token").(*jwt.Token)

		if !ok {
			return c.JSON(http.StatusInternalServerError, "Error")
		}

		c.Response().Header().Set("Authorization", "Bearer: "+token.Raw)
		return c.JSON(200, "sent header")
	}
}
