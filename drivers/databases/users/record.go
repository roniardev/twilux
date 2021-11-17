package users

import (
	"time"
	"twilux/business/users"

	"github.com/jkomyno/nanoid"
	"gorm.io/gorm"
)

type User struct {
	Id        string `gorm:"primaryKey;size:10"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Username  string         `gorm:"size:20;not null;unique"`
	Email     string         `gorm:"not null;unique"`
	Password  string         `gorm:"not null"`
}

func (user User) ToDomain() users.Domain {
	val, _ := nanoid.Nanoid(10)
	return users.Domain{
		Id:        val,
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
