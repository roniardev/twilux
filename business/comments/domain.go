package comments

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
	Comment   string `json:"comment"`
	SnippetId string `json:"snippet"`
	Snippet   snippets.Domain
	Username  string `json:"username"`
}

type CommentUsecaseInterface interface {
	GetAll(snippetId string, ctx context.Context) ([]Domain, error)
	GetAllUser(username string, ctx context.Context) ([]Domain, error)
	Create(domain Domain, ctx context.Context) (Domain, error)
	Update(domain Domain, ctx context.Context) (Domain, error)
	Delete(domain Domain, ctx context.Context) (Domain, error)
}
type CommentRepoInterface interface {
	GetAll(snippetId string, ctx context.Context) ([]Domain, error)
	GetAllUser(username string, ctx context.Context) ([]Domain, error)
	Create(domain Domain, ctx context.Context) (Domain, error)
	Update(domain Domain, ctx context.Context) (Domain, error)
	Delete(domain Domain, ctx context.Context) (Domain, error)
}
