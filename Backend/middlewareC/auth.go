package middlewareC

import (
	"fmt"
	"os"

	"github.com/Samswrld02/Portfolio/auth"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v5"
	"github.com/labstack/echo/v5"
)

// middleware for handling jwt
func CheckJwt(next echo.HandlerFunc) echo.HandlerFunc {
	config := echojwt.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
		ErrorHandler: func(c *echo.Context, err error) error {
			return c.JSON(403, map[string]string{"error: ": fmt.Errorf("not authorized").Error()})
		},
		NewClaimsFunc: func(c *echo.Context) jwt.Claims {
			return new(auth.JwtAuthClaims)
		},
	}

	middleware := echojwt.WithConfig(config)
	//run jwt middleware and handle request after
	return middleware(func(c *echo.Context) error {

		return next(c)
	})
}
