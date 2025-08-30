package api

import (
	"strings"

	"github.com/kafkaphoenix/gogrant/internal/domain"
)

// createRequestToDomain converts a request to a domain document.
func createRequestToDomain(req CreateDocumentRequest) domain.Document {
	return domain.Document{
		Title:       req.Title,
		Description: req.Description,
		Content:     req.Content,
		Tags:        req.Tags,
		Format:      req.Format,
	}
}

// createResponseFromDomain converts a domain document report to a response.
func createResponseFromDomain(report domain.DocumentReport) CreateDocumentResponse {
	return CreateDocumentResponse{
		ID:        report.ID,
		Title:     report.Title,
		CreatedAt: report.CreatedAt,
	}
}

// readResponseFromDomain converts a domain document report to a response.
func readResponseFromDomain(report domain.DocumentReport) DocumentResponse {
	return DocumentResponse{
		ID:          report.ID,
		Title:       report.Title,
		Description: report.Description,
		Content:     report.Content,
		Author:      report.Author,
		Tags:        report.Tags,
		Format:      report.Format,
		CreatedAt:   report.CreatedAt,
		UpdatedAt:   report.UpdatedAt,
	}
}

// listRequestToDomain converts a list request to a domain filter.
func listRequestToDomain(req map[string][]string) domain.DocumentFilter {
	var filter domain.DocumentFilter

	if authorVals, ok := req["author"]; ok && len(authorVals) > 0 && authorVals[0] != "" {
		filter.Author = &authorVals[0]
	}

	if tagsVals, ok := req["tags"]; ok && len(tagsVals) > 0 {
		var tags []string

		for _, raw := range tagsVals {
			// split on comma and trim spaces
			for tag := range strings.SplitSeq(raw, ",") {
				tag = strings.TrimSpace(tag)
				if tag != "" {
					tags = append(tags, tag)
				}
			}
		}

		if len(tags) > 0 {
			filter.Tags = &tags
		}
	}

	if formatVals, ok := req["format"]; ok && len(formatVals) > 0 && formatVals[0] != "" {
		filter.Format = &formatVals[0]
	}

	return filter
}

// listResponseFromDomain converts a list of domain document reports to a response.
func listResponseFromDomain(reports domain.DocumentReportList) DocumentListResponse {
	response := make(DocumentListResponse, len(reports))

	for i, report := range reports {
		response[i] = readResponseFromDomain(report)
	}

	return response
}
