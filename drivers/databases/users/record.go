package users

import (
	"time"
	"twilux/business/users"

	"gorm.io/gorm"
)

type User struct {
	Id        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Username  string
	Email     string
	Password  string
}

func (user User) ToDomain() users.Domain {
	return users.Domain{
		Id:        user.Id,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
		Username:  user.Username,
		Email:     user.Email,
		Password:  user.Password,
	}
}

func FromDomain(domain users.Domain) User {
	return User{
		Id:        domain.Id,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
		Username:  domain.Username,
		Email:     domain.Email,
		Password:  domain.Password,
	}
}
