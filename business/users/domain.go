package users

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	Username  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	Email     string
	Password  string
	Token     string
}

type UserUsecaseInterface interface {
	Login(domain Domain, ctx context.Context) (Domain, error)
	Register(domain Domain, ctx context.Context) (Domain, error)
}

type UserRepoInterface interface {
	Login(domain Domain, ctx context.Context) (Domain, error)
	Register(domain Domain, ctx context.Context) (Domain, error)
}
