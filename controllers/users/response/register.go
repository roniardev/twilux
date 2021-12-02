package response

import (
	"twilux/business/users"
	format_date "twilux/helpers/date"
)

type RegUserResponse struct {
	CreatedAt string `json:"created_at"`
	Username  string `json:"username"`
	Email     string `json:"email"`
}

func FromRegDomain(domain users.Domain) RegUserResponse {
	t := format_date.FormatDate(domain.CreatedAt.Format(format_date.LayoutISO))

	return RegUserResponse{
		CreatedAt: t,
		Username:  domain.Username,
		Email:     domain.Email,
	}
}
