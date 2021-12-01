package request

import (
	"twilux/business/saved"
)

type SavedGet struct {
	Username string `json:"snippet"`
}

func (s *SavedGet) ToGetSpecificDomain() *saved.Domain {
	return &saved.Domain{
		Username: s.Username,
	}
}
