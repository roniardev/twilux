package snippets

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	Id        string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	Title     string
	Descb     string
	Snippet   string
	Username  string
}

type SnippetUsecaseInterface interface {
	GetAll(ctx context.Context) ([]Domain, error)
	GetById(snippetId string, ctx context.Context) (Domain, error)
	Create(domain Domain, ctx context.Context) (Domain, error)
	Update(domain Domain, ctx context.Context) (Domain, error)
	Delete(domain Domain, ctx context.Context) (Domain, error)
}
type SnippetRepoInterface interface {
	GetAll(ctx context.Context) ([]Domain, error)
	GetById(snippetId string, ctx context.Context) (Domain, error)
	Create(domain Domain, ctx context.Context) (Domain, error)
	Update(domain Domain, ctx context.Context) (Domain, error)
	Delete(domain Domain, ctx context.Context) (Domain, error)
}
