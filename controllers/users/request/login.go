package request

import "twilux/business/users"

type UserLogin struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (user *UserLogin) ToDomain() *users.Domain {
	return &users.Domain{
		Email:    user.Email,
		Password: user.Password,
	}
}
