package snippets

import (
	"context"
	"twilux/helpers/models"
)

type Domain struct {
	models.Base
	Title    string `json:"title"`
	Descb    string `json:"description"`
	Snippet  string `json:"snippet"`
	Username string `json:"username"`
}

type SnippetUsecaseInterface interface {
	Create(domain Domain, ctx context.Context) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	GetByID(ctx context.Context, id uint) (Domain, error)
	Update(ctx context.Context, domain *Domain) error
	Delete(ctx context.Context, id uint) error
}
type SnippetRepoInterface interface {
	Create(domain Domain, ctx context.Context) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	GetByID(ctx context.Context, id uint) (Domain, error)
	Update(ctx context.Context, domain *Domain) error
	Delete(ctx context.Context, id uint) error
}
