package api

// swagger:schema CreateDocumentRequest
type CreateDocumentRequest struct {
	// The title of the document
	Title string `json:"title" example:"Document Title"`
	// The description of the document
	Description string `json:"description" example:"This is a sample document description."`
	// The content of the document
	Content string `json:"content" example:"This is the content of the document."`
	// The author of the document
	Author string `json:"author" example:"Author Name"`
	// The tags associated with the document
	Tags []string `json:"tags" example:"tag1,tag2"`
	// The format of the document
	Format string `json:"format" example:"markdown"`
}

// swagger:schema CreateDocumentResponse
type CreateDocumentResponse struct {
	// The ID of the created document
	ID string `json:"id" example:"6ba7b810-9dad-11d1-80b4-00c04fd430c8"`
	// The title of the created document
	Title string `json:"title" example:"Document Title"`
	// The date and time when the document was created
	CreatedAt string `json:"created_at" example:"2023-10-01T12:00:00Z"`
}

// swagger:schema ListDocumentsRequest
type ListDocumentsRequest struct {
	// Filter for documents by author
	Author string `json:"author" example:"Author Name"`
	// Filter for documents by tags
	Tags []string `json:"tags" example:"tag1,tag2"`
	// Filter for documents by format
	Format string `json:"format" example:"markdown"`
}

// swagger:schema DocumentResponse
type DocumentResponse struct {
	// The ID of the document
	ID string `json:"id" example:"6ba7b810-9dad-11d1-80b4-00c04fd430c8"`
	// The title of the document
	Title string `json:"title" example:"Document Title"`
	// The description of the document
	Description string `json:"description" example:"This is a sample document description."`
	// The content of the document
	Content string `json:"content" example:"This is the content of the document."`
	// The author of the document
	Author string `json:"author" example:"Author Name"`
	// The tags associated with the document
	Tags []string `json:"tags" example:"tag1,tag2"`
	// The format of the document
	Format string `json:"format" example:"markdown"`
	// The date and time when the document was created
	CreatedAt string `json:"created_at" example:"2023-10-01T12:00:00Z"`
	// The date and time when the document was last updated
	UpdatedAt string `json:"updated_at" example:"2023-10-01T12:00:00Z"`
}

// swagger:schema DocumentListResponse
type DocumentListResponse []DocumentResponse

// swagger:response ValidationErrorResponse
type ValidationErrorResponse struct {
	Param   string `json:"param" example:"XXX"`
	Message string `json:"message" example:"Error validating parameter"`
}

// swagger:response ErrorResponse
type ErrorResponse struct {
	Message string `json:"message" example:"Internal Server Error"`
}

// swagger:response StatusResponse
type StatusResponse struct {
	// The status of the service
	Status string `json:"status" example:"ok"`
	// The version of the service
	Version string `json:"version" example:"v0"`
}
