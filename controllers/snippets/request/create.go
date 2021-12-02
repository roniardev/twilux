package request

import (
	"twilux/business/snippets"
)

type SnippetCreate struct {
	Title    string `json:"title"`
	Descb    string `json:"descb"`
	Snippet  string `json:"snippet"`
	Username string `json:"username"`
}

func (s *SnippetCreate) ToDomain() *snippets.Domain {
	return &snippets.Domain{
		Title:    s.Title,
		Descb:    s.Descb,
		Snippet:  s.Snippet,
		Username: s.Username,
	}
}
