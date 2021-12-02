package response

import (
	"twilux/business/users"
)

type LoginUserResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

func FromLogDomain(domain users.Domain) LoginUserResponse {
	return LoginUserResponse{
		Username: domain.Username,
		Email:    domain.Email,
		Token:    domain.Token,
	}
}
