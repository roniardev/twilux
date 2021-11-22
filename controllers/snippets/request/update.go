package request

import (
	"twilux/business/snippets"
)

type SnippetUpdate struct {
	Title   string `json:"title"`
	Descb   string `json:"descb"`
	Snippet string `json:"snippet"`
}

func (s *SnippetUpdate) ToUpdateDomain() *snippets.Domain {
	return &snippets.Domain{
		Title:   s.Title,
		Descb:   s.Descb,
		Snippet: s.Snippet,
	}
}
