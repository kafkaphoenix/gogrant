package domain

import "context"

type Document struct {
	Title       string
	Description string
	Content     string
	Author      string
	Tags        []string
	Format      string
}

type DocumentReport struct {
	ID          string
	Title       string
	Description string
	Content     string
	Author      string
	Tags        []string
	Format      string
	CreatedAt   string
	UpdatedAt   string
}

type DocumentReportList []DocumentReport

type DocumentFilter struct {
	Author *string
	Tags   *[]string
	Format *string
}

type DocumentRepository interface {
	Create(ctx context.Context, document Document) (DocumentReport, error)
	Read(ctx context.Context, id string) (DocumentReport, error)
	Update(ctx context.Context, id string, document Document) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, filter DocumentFilter) (DocumentReportList, error)
}
