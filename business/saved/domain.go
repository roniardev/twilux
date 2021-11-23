package saved

import (
	"context"
	"time"
	"twilux/business/snippets"

	"gorm.io/gorm"
)

type Domain struct {
	Id        string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	SnippetId string `json:"snippet"`
	Snippet   snippets.Domain
	Username  string `json:"username"`
}

type SavedUsecaseInterface interface {
	GetAll(ctx context.Context) ([]Domain, error)
	Create(domain Domain, ctx context.Context) (Domain, error)
	Delete(domain Domain, ctx context.Context) (Domain, error)
}

type SavedRepoInterface interface {
	GetAll(ctx context.Context) ([]Domain, error)
	Create(domain Domain, ctx context.Context) (Domain, error)
	Delete(domain Domain, ctx context.Context) (Domain, error)
}
