package snippets

import (
	"context"
	"fmt"
	"twilux/business/snippets"

	"github.com/jkomyno/nanoid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	errs := repo.db.Preload(clause.Associations).Preload("Snippet."+clause.Associations).First(&snippetDb, snippetDb).Error
	if errs != nil {
		return snippetDb.ToDomain(), nil
	}
	return snippetDb.ToDomain(), nil
}

// Update specific snippet by id
func (repo *SnippetRepository) Update(domain snippets.Domain, ctx context.Context) (snippets.Domain, error) {
	snippetDb := FromDomain(domain)
	res := repo.db.Updates(&snippetDb)
	if res.Error != nil {
		return snippets.Domain{}, res.Error
	}

	return snippetDb.ToDomain(), nil
}

// Update deleted_at field to specific snippet by id
func (repo *SnippetRepository) Delete(domain snippets.Domain, ctx context.Context) (snippets.Domain, error) {
	snippetDb := FromDomain(domain)
	res := repo.db.Where("username = ?", snippetDb.Username).Delete(&snippetDb)
	if res.Error != nil {
		return snippets.Domain{}, res.Error
	}

	return snippetDb.ToDomain(), nil
}
