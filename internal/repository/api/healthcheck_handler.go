package api

import (
	"log/slog"
	"net/http"
)

// HealthHandler is a struct that implements the HTTP handler for health check requests.
type HealthHandler struct {
	logger *slog.Logger
}

// NewHealthHandler is a factory to instantiate a new HealthHandler.
func NewHealthHandler(logger *slog.Logger) *HealthHandler {
	return &HealthHandler{
		logger: logger.With("component", "health_handler"),
	}
}

// RegisterRoutes register the handler routes.
func (h *HealthHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/v0/health", h.health)
}

// @Summary      Health check
// @Description  Returns the health status of the service.
// @ID           health
// @Tags         health
// @Produce      json
// @Success      200  {object} StatusResponse
// @Router       /api/v0/health [get]
//
// health is the handler function for the health check endpoint.
func (h *HealthHandler) health(w http.ResponseWriter, _ *http.Request) {
	WriteResponse(h.logger, w, http.StatusOK, StatusResponse{Status: "ok", Version: "v0"})
}
