package api

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/google/uuid"
	"github.com/kafkaphoenix/gogrant/internal/usecase"
)

// DocumentHandler is a struct that implements the HTTP handler for document requests.
type DocumentHandler struct {
	logger  *slog.Logger
	service *usecase.DocumentService
}

// NewDocumentHandler is a factory to instantiate a new DocumentHandler.
func NewDocumentHandler(logger *slog.Logger, service *usecase.DocumentService) *DocumentHandler {
	return &DocumentHandler{
		logger:  logger.With("component", "document_handler"),
		service: service,
	}
}

// RegisterRoutes register the handler routes.
func (h *DocumentHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/v0/documents", h.createDocument)
	mux.HandleFunc("GET /api/v0/documents/{id}", h.readDocument)
	mux.HandleFunc("GET /api/v0/documents", h.listDocuments)
}

// @Summary      Create document
// @Description  Creates a new document.
// @ID           create-document
// @Tags         documents
// @Accept       json
// @Produce      json
// @Param        document body CreateDocumentRequest true "Document to create"
// @Success      201  {object} CreateDocumentResponse
// @Failure      500  {object} ErrorResponse
// @Router       /api/v0/documents [post]
//
// createDocument handles the request to create a new document.
func (h *DocumentHandler) createDocument(w http.ResponseWriter, r *http.Request) {
	defer CloseBody(h.logger, r.Body)

	var req CreateDocumentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		WriteValidationError(h.logger, w, "body", "Invalid request body: "+err.Error())
		return
	}

	// TODO validate the request

	document := createRequestToDomain(req)

	report, err := h.service.CreateDocument(r.Context(), document)
	if err != nil {
		WriteError(h.logger, w, http.StatusInternalServerError, err.Error())
		return
	}

	response := createResponseFromDomain(report)

	WriteResponse(h.logger, w, http.StatusCreated, response)
}

// @Summary      Read document
// @Description  Returns a document by its ID.
// @ID           read-document
// @Tags         documents
// @Produce      json
// @Param        id path string true "Document ID" example("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
// @Success      200  {object} DocumentResponse
// @Failure      400 {object} ValidationErrorResponse "Device ID is required"
// @Failure      400 {object} ValidationErrorResponse "Invalid device ID format"
// @Failure      500  {object} ErrorResponse
// @Router       /api/v0/documents/{id} [get]
func (h *DocumentHandler) readDocument(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		WriteValidationError(h.logger, w, "id", "Document ID is required")
		return
	}

	_, err := uuid.Parse(id)
	if err != nil {
		WriteValidationError(h.logger, w, "id", "Invalid document ID format")
		return
	}

	report, err := h.service.ReadDocument(r.Context(), id)
	if err != nil {
		WriteError(h.logger, w, http.StatusInternalServerError, err.Error())
		return
	}

	response := readResponseFromDomain(report)

	WriteResponse(h.logger, w, http.StatusOK, response)
}

// @Summary      List documents
// @Description  Returns a list of documents based on the provided filter.
// @ID           list-documents
// @Tags         documents
// @Produce      json
// @Param        author query string false "Filter by author" example("Author Name")
// @Param        tags query string false "Filter by tags" example("tag1,tag2")
// @Param        format query string false "Filter by format" example("markdown")
// @Success      200  {object} DocumentListResponse
// @Failure      500  {object} ErrorResponse
// @Router       /api/v0/documents [get]
//
// listDocuments handles the request to list documents with an optional filter.
func (h *DocumentHandler) listDocuments(w http.ResponseWriter, r *http.Request) {
	filter := listRequestToDomain(r.URL.Query())

	documents, err := h.service.ListDocuments(r.Context(), filter)
	if err != nil {
		WriteError(h.logger, w, http.StatusInternalServerError, err.Error())
		return
	}

	response := listResponseFromDomain(documents)

	WriteResponse(h.logger, w, http.StatusOK, response)
}
