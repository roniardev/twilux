package snippets

import (
	"context"
	"time"
	"twilux/business/users"

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
	UserInfo  users.Domain
}

type SnippetUsecaseInterface interface {
	GetAll(ctx context.Context) ([]Domain, error)
	Create(domain Domain, ctx context.Context) (Domain, error)
	Update(domain Domain, ctx context.Context) (Domain, error)
}
type SnippetRepoInterface interface {
	GetAll(ctx context.Context) ([]Domain, error)
	Create(domain Domain, ctx context.Context) (Domain, error)
	Update(domain Domain, ctx context.Context) (Domain, error)
}
