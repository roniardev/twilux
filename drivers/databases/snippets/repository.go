package snippets

import (
	"context"
	"twilux/business/snippets"

	"gorm.io/gorm"
)

type SnippetRepository struct {
	db *gorm.DB
}

func NewSnippetRepository(db *gorm.DB) *SnippetRepository {
	return &SnippetRepository{db}
}

func (repo *SnippetRepository) GetAllSnippets(ctx context.Context) ([]snippets.Domain, error) {
	return []snippets.Domain{}, nil
}
