// Package api provides the HTTP server and route registration for the signing service API.
package api

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	httpSwagger "github.com/swaggo/http-swagger/v2"
)

const (
	APIVersion  = "v0"
	APIBasePath = "/api/" + APIVersion
)

// RouteRegister represents a function to register routes in the router.
type RouteRegister func(mux *http.ServeMux)

// Server contains all the data related with the http server.
type Server struct {
	logger *slog.Logger
	srv    *http.Server
	port   int
}

// NewServer returns a new instance of Server struct.
func NewServer(logger *slog.Logger, port int) (*Server, error) {
	if port <= 0 {
		return nil, ErrInvalidPort
	}

	server := &Server{
		logger: logger.With("component", "api"),
		srv: &http.Server{
			Addr:         ":" + strconv.Itoa(port),
			Handler:      http.NewServeMux(),
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
		port: port,
	}

	return server, nil
}

// RegisterRoutes given n handlers will register all of them.
func (s *Server) RegisterRoutes(handlers ...RouteRegister) error {
	router, ok := s.srv.Handler.(*http.ServeMux)
	if !ok {
		return &RouterError{}
	}

	// register swagger
	router.HandleFunc(APIBasePath+"/swagger/", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:"+strconv.Itoa(s.port)+APIBasePath+"/swagger/doc.json"),
	))

	for _, handler := range handlers {
		handler(router)
	}

	s.logger.Info("Routes registered")

	return nil
}

// Start creates a new server with graceful shutdown.
func (s *Server) Start(ctx context.Context) error {
	quit, stop := signal.NotifyContext(ctx, os.Interrupt, os.Kill)
	defer stop()

	s.logger.Info("Starting server...", "port", s.port)

	go func() {
		if err := s.srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			s.logger.Error("server error", "error", err.Error())
		}

		stop() // in case server returns before signal is received
	}()

	<-quit.Done()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	s.logger.InfoContext(ctx, "Shutdown signal received")

	if err := s.Shutdown(ctx); err != nil {
		return &ServerError{
			Message: "error shutting down server",
			Err:     err,
		}
	}

	s.logger.InfoContext(ctx, "Server gracefully stopped")

	return nil
}

// Shutdown closes the server.
func (s *Server) Shutdown(ctx context.Context) error {
	s.logger.InfoContext(ctx, "Shutting down server...")
	return s.srv.Shutdown(ctx)
}

// WriteResponse writes a JSON response to the http.ResponseWriter.
func WriteResponse(logger *slog.Logger, w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(payload); err != nil {
		WriteError(logger, w, http.StatusInternalServerError, "failed to encode response: "+err.Error())
		return
	}
}

// WriteValidationError writes a validation error response to the http.ResponseWriter.
func WriteValidationError(logger *slog.Logger, w http.ResponseWriter, param, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)

	err := json.NewEncoder(w).Encode(ValidationErrorResponse{
		Param:   param,
		Message: message,
	})
	if err != nil {
		logger.Error("Failed to encode validation error response", "error", err)
		return
	}
}

// WriteError writes an error response to the http.ResponseWriter.
func WriteError(logger *slog.Logger, w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	err := json.NewEncoder(w).Encode(ErrorResponse{
		Message: message,
	})
	if err != nil {
		logger.Error("Failed to encode error response", "error", err)
		return
	}
}

// CloseBody drains and closes the request body to avoid resource leaks.
func CloseBody(logger *slog.Logger, body io.ReadCloser) {
	// drain anything left in the body and close it, to avoid resource leaks
	if _, err := io.Copy(io.Discard, body); err != nil {
		logger.Error("Error discarding body", "error", err)
	}

	if err := body.Close(); err != nil {
		logger.Error("Error closing body", "error", err)
	}
}
