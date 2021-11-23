package request

import (
	"twilux/business/saved"
)

type SavedCreate struct {
	SnippetId string `json:"snippetId"`
	Username  string `json:"username"`
}

func (s *SavedCreate) ToDomain() *saved.Domain {
	return &saved.Domain{
		SnippetId: s.SnippetId,
		Username:  s.Username,
	}
}
