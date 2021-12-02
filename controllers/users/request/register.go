package request

import "twilux/business/users"

type UserRegister struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

func (u *UserRegister) ToDomain() *users.Domain {
	return &users.Domain{
		Email:    u.Email,
		Password: u.Password,
		Username: u.Username,
	}
}
