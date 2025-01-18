package main

import (
	"log"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/satriadhm-echo-boilerplate/internal/auth/delivery/http"
	authRepo "github.com/satriadhm-echo-boilerplate/internal/auth/repository"
	authUsecase "github.com/satriadhm-echo-boilerplate/internal/auth/usecase"
	"github.com/satriadhm-echo-boilerplate/internal/middleware"
	todoHttp "github.com/satriadhm-echo-boilerplate/internal/todo/delivery/http"
	todoRepo "github.com/satriadhm-echo-boilerplate/internal/todo/repository"
	todoUsecase "github.com/satriadhm-echo-boilerplate/internal/todo/usecase"
	"github.com/satriadhm-echo-boilerplate/pkg/config"
	"github.com/satriadhm-echo-boilerplate/pkg/logger"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig("configs/config.yaml")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize logger
	logger.Init(cfg.Logging.Level, cfg.Logging.File)

	// Create Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.ErrorHandler)
	if cfg.Middlewares.EnableRequestLogging {
		e.Use(middleware.RequestLogger)
	}

	// Database connection
	db, err := config.ConnectDatabase(cfg.Database)
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
	http.NewAuthHandler(e, authUC)
	todoHttp.NewTodoHandler(e, todoUC)

	// Start the server
	serverConfig := middleware.TimeoutMiddleware(cfg.Server.ReadTimeout, cfg.Server.WriteTimeout)
	e.Server.ReadTimeout, _ = time.ParseDuration(cfg.Server.ReadTimeout)
	e.Server.WriteTimeout, _ = time.ParseDuration(cfg.Server.WriteTimeout)

	e.Logger.Fatal(e.Start(cfg.Server.Port))
}
