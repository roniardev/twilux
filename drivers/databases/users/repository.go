package users

import (
	"context"
	"twilux/business/users"

	"github.com/jkomyno/nanoid"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(gormDb *gorm.DB) users.UserRepoInterface {
	return &UserRepository{
		db: gormDb,
	}
}

func (repo *UserRepository) Login(domain users.Domain, ctx context.Context) (users.Domain, error) {
	userDb := FromDomain(domain)

	err := repo.db.Where("email = ? AND password = ?", userDb.Email, userDb.Password).Error
	if err != nil {
		return users.Domain{}, err
	}
	return userDb.ToDomain(), nil
}

// SignUp creates a new user
func (repo *UserRepository) Register(domain users.Domain, ctx context.Context) (users.Domain, error) {
	userDb := FromDomain(domain)
	userDb.Id, _ = nanoid.Nanoid(10)

	err := repo.db.Create(&userDb).Error
	if err != nil {
		return users.Domain{}, err
	}
	return userDb.ToDomain(), nil
}
