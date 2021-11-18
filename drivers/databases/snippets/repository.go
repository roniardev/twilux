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

func (repo *SnippetRepository) GetAll(ctx context.Context) ([]snippets.Domain, error) {
	return []snippets.Domain{}, nil
}

func (repo *SnippetRepository) Create(domain snippets.Domain, ctx context.Context) (snippets.Domain, error) {
	snippetDb := FromDomain(domain)

	err := repo.db.Create(&snippetDb).Error
	if err != nil {
		return snippets.Domain{}, err
	}
	return snippetDb.ToDomain(), nil
}
