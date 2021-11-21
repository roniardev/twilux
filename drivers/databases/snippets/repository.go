package snippets

import (
	"context"
	"fmt"
	"twilux/business/snippets"

	"github.com/jkomyno/nanoid"
	"gorm.io/gorm"
)

type SnippetRepository struct {
	db *gorm.DB
}

func NewSnippetRepository(db *gorm.DB) *SnippetRepository {
	return &SnippetRepository{db}
}

func (repo *SnippetRepository) GetAll(ctx context.Context) ([]snippets.Domain, error) {
	snipp := []Snippet{}
	result := repo.db.Find(&snipp)
	if result.Error != nil {
		return []snippets.Domain{}, result.Error
	}
	fmt.Println("GetAll Repo db")
	fmt.Println(result)
	return ToListDomain(snipp), nil

}

func (repo *SnippetRepository) Create(domain snippets.Domain, ctx context.Context) (snippets.Domain, error) {
	snippetDb := FromDomain(domain)
	snippetDb.Id, _ = nanoid.Nanoid(10)

	err := repo.db.Create(&snippetDb).Error
	if err != nil {
		return snippets.Domain{}, err
	}
	return snippetDb.ToDomain(), nil
}
