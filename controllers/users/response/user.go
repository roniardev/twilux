package response

import (
	"time"
	"twilux/business/users"

	"gorm.io/gorm"
)

type UserResponse struct {
	Id        string         `json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
	Email     string         `json:"email"`
	Username  string         `json:"username"`
	Password  string         `json:"password"`
	Token     string         `json:"token"`
}

func FromDomain(domain users.Domain) UserResponse {
	return UserResponse{
		Id:        domain.Id,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
		Username:  domain.Username,
		Email:     domain.Email,
		Password:  domain.Password,
		Token:     domain.Token,
	}
}
