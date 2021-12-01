package response

import (
	"time"
	"twilux/business/comments"
	resSnip "twilux/controllers/snippets/response"
)

const (
	layoutISO = "2006-01-02T15:04:05.999999999Z07:00"
	layoutUS  = "Monday, January 2, 2006, 03:04:01 PM"
)

type CommentResponse struct {
	Id        string                        `json:"id"`
	Username  string                        `json:"username"`
	CreatedAt string                        `json:"created_at"`
	Comment   string                        `json:"comment"`
	SnippetId string                        `json:"snippet_id"`
	Snippet   resSnip.SnippetCreateResponse `json:"snippet"`
}

func FromDomain(domain comments.Domain) CommentResponse {
	t, _ := time.Parse(layoutISO, domain.CreatedAt.Format(layoutISO))

	return CommentResponse{
		CreatedAt: t.Format(layoutUS),
		Id:        domain.Id,
		Username:  domain.Username,
		Comment:   domain.Comment,
		SnippetId: domain.SnippetId,
		Snippet:   resSnip.FromCreateDomain(domain.Snippet),
	}
}
