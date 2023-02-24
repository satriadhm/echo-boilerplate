package main

import (
	"github.com/Satria2133/echo-CRUD/auth"
	"github.com/Satria2133/echo-CRUD/config"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Login route
	e.POST("/login", auth.Login)

	// Unauthenticated route
	e.GET("/", auth.Accessible)

	// Restricted group
	r := e.Group("/restricted")

	r.Use(echojwt.WithConfig(config.AuthConfig()))

	r.GET("", auth.Restricted)

	e.Logger.Fatal(e.Start(":1323"))
}
