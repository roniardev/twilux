package saved

import (
	"context"
	"fmt"
	"twilux/business/saved"

	"github.com/jkomyno/nanoid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SavedRepository struct {
	db *gorm.DB
}

func NewSavedRepository(db *gorm.DB) *SavedRepository {
	return &SavedRepository{db}
}

func (repo *SavedRepository) GetAll(ctx context.Context) ([]saved.Domain, error) {
	sav := []Saved{}
	result := repo.db.Find(&sav)
	if result.Error != nil {
		return []saved.Domain{}, result.Error
	}
	fmt.Println("GetAll Repo db")
	fmt.Println(result)
	return ToListDomain(sav), nil

}

func (repo *SavedRepository) Create(domain saved.Domain, ctx context.Context) (saved.Domain, error) {
	savedDb := FromDomain(domain)
	savedDb.Id, _ = nanoid.Nanoid(10)

	err := repo.db.Create(&savedDb).Error
	if err != nil {
		return saved.Domain{}, err
	}
	errs := repo.db.Preload(clause.Associations).Preload("Saved."+clause.Associations).First(&savedDb, savedDb).Error
	if errs != nil {
		return savedDb.ToDomain(), nil
	}
	return savedDb.ToDomain(), nil
}

// Update deleted_at field to specific snippet by id
func (repo *SavedRepository) Delete(domain saved.Domain, ctx context.Context) (saved.Domain, error) {
	savedDb := FromDomain(domain)
	res := repo.db.Where("id = ?", savedDb.Username).Delete(&savedDb)
	if res.Error != nil {
		return saved.Domain{}, res.Error
	}

	return savedDb.ToDomain(), nil
}
