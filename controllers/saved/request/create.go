package request

import (
	"twilux/business/saved"
)

type SavedCreate struct {
	SnippetId string `json:"snippet_id"`
	Username  string `json:"username"`
}

func (s *SavedCreate) ToDomain() *saved.Domain {
	return &saved.Domain{
		SnippetId: s.SnippetId,
		Username:  s.Username,
	}
}
