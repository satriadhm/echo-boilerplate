package config

import (
	"github.com/Satria2133/echo-CRUD/auth"
	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

// Configure middleware with the custom claims type

func AuthConfig() echojwt.Config {
	Config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(auth.JwtCustomClaims)
		},
		SigningKey: []byte("secret"),
	}
	return Config
}
