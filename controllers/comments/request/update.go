package request

import (
	"twilux/business/comments"
)

type CommentUpdate struct {
	Id       string `json:"id"`
	Comment  string `json:"comment"`
	Username string `json:"username"`
}

func (s *CommentUpdate) ToUpdateDomain() *comments.Domain {
	return &comments.Domain{
		Id:       s.Id,
		Comment:  s.Comment,
		Username: s.Username,
	}
}
