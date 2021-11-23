package request

import (
	"twilux/business/saved"
)

type SavedDelete struct {
	Id       string `json:"id"`
	Username string `json:"snippet"`
}

func (s *SavedDelete) ToDeleteDomain() *saved.Domain {
	return &saved.Domain{
		Id:       s.Id,
		Username: s.Username,
	}
}
