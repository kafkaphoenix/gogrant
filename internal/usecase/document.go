package usecase

import (
	"context"

	"github.com/kafkaphoenix/gogrant/internal/domain"
)

type DocumentService struct {
	documentRepository domain.DocumentRepository
}

func NewDocumentService(documentRepository domain.DocumentRepository) *DocumentService {
	return &DocumentService{
		documentRepository: documentRepository,
	}
}
func (s *DocumentService) CreateDocument(ctx context.Context, document domain.Document) (domain.DocumentReport, error) {
	return s.documentRepository.Create(ctx, document)
}

func (s *DocumentService) ReadDocument(ctx context.Context, id string) (domain.DocumentReport, error) {
	return s.documentRepository.Read(ctx, id)
}

func (s *DocumentService) UpdateDocument(ctx context.Context, id string, document domain.Document) error {
	return s.documentRepository.Update(ctx, id, document)
}

func (s *DocumentService) DeleteDocument(ctx context.Context, id string) error {
	return s.documentRepository.Delete(ctx, id)
}

func (s *DocumentService) ListDocuments(ctx context.Context, filter domain.DocumentFilter) (domain.DocumentReportList, error) {
	return s.documentRepository.List(ctx, filter)
}
