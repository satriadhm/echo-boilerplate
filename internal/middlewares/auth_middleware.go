package middleware

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// JWTAuthMiddleware configures JWT middleware
func JWTAuthMiddleware() echo.MiddlewareFunc {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		panic("JWT_SECRET environment variable is required")
	}

	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(secret),
	})
}
