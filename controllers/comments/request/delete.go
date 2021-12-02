package request

import (
	"twilux/business/comments"
)

type CommentDelete struct {
	Id        string `json:"id"`
	Username  string `json:"username"`
	SnippetId string `json:"snippet_id"`
}

func (s *CommentDelete) ToDeleteDomain() *comments.Domain {
	return &comments.Domain{
		Id:        s.Id,
		Username:  s.Username,
		SnippetId: s.SnippetId,
	}
}
