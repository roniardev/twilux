package request

import (
	"twilux/business/saved"
)

type SavedDelete struct {
	SnippetId string `json:"snippet_id"`
	Username  string `json:"username"`
}

func (s *SavedDelete) ToDeleteDomain() *saved.Domain {
	return &saved.Domain{
		SnippetId: s.SnippetId,
		Username:  s.Username,
	}
}
