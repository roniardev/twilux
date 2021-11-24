package request

import (
	"twilux/business/comments"
)

type CommentCreate struct {
	SnippetId string `json:"snippetId"`
	Comment   string `json:"comment"`
	Username  string `json:"username"`
}

func (s *CommentCreate) ToDomain() *comments.Domain {
	return &comments.Domain{
		SnippetId: s.SnippetId,
		Comment:   s.Comment,
		Username:  s.Username,
	}
}
