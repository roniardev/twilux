package saved

import (
	"context"
	"time"
	"twilux/business/snippets"
	"twilux/business/users"

	"gorm.io/gorm"
)

type Domain struct {
	Id        string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	SnippetId string
	Snippet   snippets.Domain
	Username  string
	UserInfo  users.Domain
}

type SavedUsecaseInterface interface {
	GetAll(username string, ctx context.Context) ([]Domain, error)
	Create(domain Domain, ctx context.Context) (Domain, error)
	Delete(domain Domain, ctx context.Context) (Domain, error)
}

type SavedRepoInterface interface {
	GetAll(username string, ctx context.Context) ([]Domain, error)
	Create(domain Domain, ctx context.Context) (Domain, error)
	Delete(domain Domain, ctx context.Context) (Domain, error)
}
