package response

import (
	"time"
	"twilux/business/users"
)

type RegUserResponse struct {
	CreatedAt time.Time `json:"created_at"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
}

func FromRegDomain(domain users.Domain) RegUserResponse {
	return RegUserResponse{
		CreatedAt: domain.CreatedAt,
		Username:  domain.Username,
		Email:     domain.Email,
		Password:  domain.Password,
	}
}
