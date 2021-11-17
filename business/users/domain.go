package users

import (
	"context"
	"twilux/helpers/models"
)

type Domain struct {
	models.Base
	Email    string
	Username string
	Password string
	Token    string
}

type UserUsecaseInterface interface {
	Login(domain Domain, ctx context.Context) (Domain, error)
	GetAllUsers(ctx context.Context) ([]Domain, error)
}

type UserRepoInterface interface {
	Login(domain Domain, ctx context.Context) (Domain, error)
	GetAllUsers(ctx context.Context) ([]Domain, error)
}
