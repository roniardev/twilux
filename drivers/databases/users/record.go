package users

import (
	"time"
	"twilux/business/users"

	"gorm.io/gorm"
)

type User struct {
	Username  string `gorm:"size:20;not null;unique;primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Email     string         `gorm:"not null;unique"`
	Password  string         `gorm:"not null"`
}

func (user User) ToDomain() users.Domain {

	return users.Domain{
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
		Email:     user.Email,
		Password:  user.Password,
	}
}

func FromDomain(domain users.Domain) User {
	return User{
		Username:  domain.Username,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
		Email:     domain.Email,
		Password:  domain.Password,
	}
}
