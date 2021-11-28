package request

import (
	"twilux/business/snippets"
)

type SnippetUpdate struct {
	Id       string `json:"id"`
	Title    string `json:"title"`
	Descb    string `json:"descb"`
	Snippet  string `json:"snippet"`
	Username string `json:"usrname"`
}

func (s *SnippetUpdate) ToUpdateDomain() *snippets.Domain {
	return &snippets.Domain{
		Id:       s.Id,
		Title:    s.Title,
		Descb:    s.Descb,
		Snippet:  s.Snippet,
		Username: s.Username,
	}
}
