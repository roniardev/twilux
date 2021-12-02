package saved

import (
	"context"
	"errors"
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

func (repo *SavedRepository) GetAll(user string, ctx context.Context) ([]saved.Domain, error) {
	sav := []Saved{}
	result := repo.db.Preload(clause.Associations).Preload("Snippet."+clause.Associations).Where("user = ?", user).Find(&sav)

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

	isDuplicate := repo.db.First(&Saved{}, "user = ? AND snippet_id = ?", savedDb.User, savedDb.SnippetId)
	if isDuplicate.Error == nil {
		return saved.Domain{}, errors.New("duplicate snippet saved")
	}

	err := repo.db.Create(&savedDb).Error
	if err != nil {
		return saved.Domain{}, err
	}
	errs := repo.db.Preload(clause.Associations).Preload("Snippet." + clause.Associations).Find(&savedDb)

	if errs != nil {
		return savedDb.ToDomain(), nil
	}
	return savedDb.ToDomain(), nil
}

// Update deleted_at field to specific snippet by id
func (repo *SavedRepository) Delete(domain saved.Domain, ctx context.Context) (saved.Domain, error) {
	savedDb := FromDomain(domain)
	isEligible := repo.db.Where("user = ? AND snippet_id = ?", savedDb.User, savedDb.SnippetId).First(&savedDb, savedDb)

	if isEligible.Error != nil {
		return saved.Domain{}, errors.New("you are not eligible to delete this saved")
	}

	res := repo.db.Where("user = ?", savedDb.User).Delete(&savedDb)
	if res.Error != nil {
		return saved.Domain{}, res.Error
	}

	return savedDb.ToDomain(), nil
}
