package users

import (
	"context"
	"twilux/business/users"

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

	err := repo.db.First(&userDb, "email = ? ", userDb.Email).Error
	if err != nil {
		return users.Domain{}, err
	}
	return userDb.ToDomain(), nil
}

// SignUp creates a new user
func (repo *UserRepository) Register(domain users.Domain, ctx context.Context) (users.Domain, error) {
	userDb := FromDomain(domain)

	err := repo.db.Create(&userDb).Error
	if err != nil {
		return users.Domain{}, err
	}
	return userDb.ToDomain(), nil
}
