package main

import (
	"github.com/Satria2133/echo-CRUD/auth"
	"github.com/Satria2133/echo-CRUD/config"
	controller "github.com/Satria2133/echo-CRUD/controller"
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

	//routing CRUD
	e.POST("/todo", controller.NewTodoList)
	e.GET("/todo/:id", controller.GetTodoList)
	e.PUT("/todo/:id", controller.UpdateTodoList)
	e.DELETE("/todo/:id", controller.DeleteTodoList)
	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
