package mongodb

import (
	"context"
	"log/slog"

	"github.com/kafkaphoenix/gogrant/internal/domain"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

const _documentCollectionName = "documents"

type MongoDocumentRepository struct {
	logger     *slog.Logger
	collection *mongo.Collection
}

// interface guard
var _ domain.DocumentRepository = (*MongoDocumentRepository)(nil)

func NewDocumentRepository(logger *slog.Logger, db *MongoDB) *MongoDocumentRepository {
	return &MongoDocumentRepository{
		logger:     logger.With("component", "document_repository"),
		collection: db.Collection(_documentCollectionName),
	}
}

func (r *MongoDocumentRepository) Create(ctx context.Context, document domain.Document) (domain.DocumentReport, error) {
	r.logger.Debug("creating document", "document", document)

	return domain.DocumentReport{
		Title:       document.Title,
		Description: document.Description,
		Content:     document.Content,
	}, nil
}

func (r *MongoDocumentRepository) Read(ctx context.Context, id string) (domain.DocumentReport, error) {
	r.logger.Debug("reading document", "id", id)

	return domain.DocumentReport{
		ID:          id,
		Title:       "Sample Title",
		Description: "Sample Description",
		Content:     "Sample Content",
		Author:      "Sample Author",
		Tags:        []string{"tag1", "tag2"},
		Format:      "markdown",
		CreatedAt:   "2023-10-01T00:00:00Z",
		UpdatedAt:   "2023-10-01T00:00:00Z",
	}, nil
}

func (r *MongoDocumentRepository) Update(ctx context.Context, id string, document domain.Document) error {
	r.logger.Debug("updating document", "id", id)
	return nil
}

func (r *MongoDocumentRepository) Delete(ctx context.Context, id string) error {
	r.logger.Debug("deleting document", "id", id)
	return nil
}

func (r *MongoDocumentRepository) List(ctx context.Context, filter domain.DocumentFilter) (domain.DocumentReportList, error) {
	r.logger.Debug("listing documents", "filter", filter)

	return domain.DocumentReportList{
		{
			ID:          "6ba7b810-9dad-11d1-80b4-00c04fd430c8",
			Title:       "Sample Title",
			Description: "Sample Description",
			Content:     "Sample Content",
			Author:      "Sample Author",
			Tags:        []string{"tag1", "tag2"},
			Format:      "markdown",
			CreatedAt:   "2023-10-01T00:00:00Z",
			UpdatedAt:   "2023-10-01T00:00:00Z",
		},
		{
			ID:          "6ba7b810-9dad-11d1-80b4-00c04fd430c9",
			Title:       "Sample Title 2",
			Description: "Sample Description 2",
			Content:     "Sample Content 2",
			Author:      "Sample Author",
			Tags:        []string{"tag1", "tag2"},
			Format:      "markdown",
			CreatedAt:   "2023-10-02T00:00:00Z",
			UpdatedAt:   "2023-10-02T00:00:00Z",
		},
	}, nil
}
