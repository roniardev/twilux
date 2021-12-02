package request

import (
	"twilux/business/snippets"
)

type SnippetDelete struct {
	Id       string `json:"id"`
	Username string `json:"snippet"`
}

func (s *SnippetDelete) ToDeleteDomain() *snippets.Domain {
	return &snippets.Domain{
		Id:       s.Id,
		Username: s.Username,
	}
}
