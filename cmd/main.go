package main

import (
	"log"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	authHttp "github.com/satriadhm/echo-boilerplate/internal/auth/delivery/http"
	authRepo "github.com/satriadhm/echo-boilerplate/internal/auth/repository"
	authUsecase "github.com/satriadhm/echo-boilerplate/internal/auth/usecase"
	todoHttp "github.com/satriadhm/echo-boilerplate/internal/todo/delivery/http"
	todoRepo "github.com/satriadhm/echo-boilerplate/internal/todo/repository"
	todoUsecase "github.com/satriadhm/echo-boilerplate/internal/todo/usecase"
	"github.com/satriadhm/echo-boilerplate/pkg/config"
	"github.com/satriadhm/echo-boilerplate/pkg/logger"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig("config.yaml.example")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize logger
	logger.Init(cfg.Logging.Level, cfg.Logging.File)

	// Create Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	if cfg.Middlewares.EnableRequestLogging {
		e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
			LogURI:    true,
			LogStatus: true,
			LogError:  true,
		}))
	}

	db, err := config.ConnectDatabase(*cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Repository and Usecase setup
	authRepository := authRepo.NewAuthRepository(db)
	authUC := authUsecase.NewAuthUsecase(authRepository)

	todoRepository := todoRepo.NewTodoRepository(db)
	todoUC := todoUsecase.NewTodoUsecase(todoRepository)

	// Routes setup
	authHttp.NewAuthHandler(e, authUC)
	todoHttp.NewTodoHandler(e, todoUC)

	// Server configuration
	readTimeout, _ := time.ParseDuration(cfg.Server.ReadTimeout)
	writeTimeout, _ := time.ParseDuration(cfg.Server.WriteTimeout)
	e.Server.ReadTimeout = readTimeout
	e.Server.WriteTimeout = writeTimeout

	// Start the server
	e.Logger.Fatal(e.Start(cfg.Server.Port))
}
