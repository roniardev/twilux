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
	Title     string `json:"title"`
	Descb     string `json:"description"`
	Snippet   string `json:"snippet"`
	Username  string `json:"username"`
}

type SnippetUsecaseInterface interface {
	Create(domain Domain, ctx context.Context) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
}
type SnippetRepoInterface interface {
	Create(domain Domain, ctx context.Context) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
}
