//go:generate go tool mockery
package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"strings"

	_ "github.com/kafkaphoenix/gogrant/docs"
	"github.com/kafkaphoenix/gogrant/internal/repository/api"
	"github.com/kafkaphoenix/gogrant/internal/repository/config"
	"github.com/kafkaphoenix/gogrant/internal/repository/mongodb"
	"github.com/kafkaphoenix/gogrant/internal/usecase"
)

// ServiceError represents service errors.
type ServiceError struct {
	Message string
	Err     error
}

func (e *ServiceError) Error() string {
	return fmt.Sprintf("%s: %v", e.Message, e.Err)
}

func (e *ServiceError) Unwrap() error {
	return e.Err
}

// run starts the service by creating the logger, external services, usecases, handlers and the http server.
func run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger := initLogger()
	slog.SetDefault(logger)
	logger.Info("Starting Grant Service...")

	cfg, err := config.Load(logger)
	if err != nil {
		return &ServiceError{Message: "failed to load config", Err: err}
	}

	db := mongodb.NewMongoDB(logger)
	if err = db.Connect(ctx, cfg.Mongo.URI); err != nil {
		return &ServiceError{Message: "failed to start database", Err: err}
	}

	defer db.Disconnect(ctx)

	health := api.NewHealthHandler(logger)
	documentRepo := mongodb.NewDocumentRepository(logger, db)
	documentService := usecase.NewDocumentService(documentRepo)
	documentHandler := api.NewDocumentHandler(logger, documentService)

	srv, err := api.NewServer(logger, cfg.App.Port)
	if err != nil {
		return &ServiceError{Message: "failed to create server", Err: err}
	}

	if err = srv.RegisterRoutes(health.RegisterRoutes, documentHandler.RegisterRoutes); err != nil {
		return &ServiceError{Message: "failed to register routes", Err: err}
	}

	if err = srv.Start(ctx); err != nil {
		return &ServiceError{Message: "failed to start server", Err: err}
	}

	return nil
}

// initLogger initializes the structured logger based on the DEBUG_ENABLED environment variable.
func initLogger() *slog.Logger {
	debugEnabled := strings.ToLower(os.Getenv("DEBUG_ENABLED")) == "true"

	level := slog.LevelInfo
	if debugEnabled {
		level = slog.LevelDebug
	}

	return slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: level}))
}

// @title Grant Service API
// @version 1.0
// @description This is the API documentation for the Grant Service.

// @contact.name Javier Aguilera
// @contact.email jaguilerapuerta@gmail.com

// @schemes http
// @host localhost:8080
// @BasePath /
func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
