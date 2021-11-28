package response

import (
	"time"
	"twilux/business/comments"
	resSnip "twilux/controllers/snippets/response"
)

type CommentUpdateResponse struct {
	Id        string                        `json:"id"`
	Username  string                        `json:"username"`
	UpdatedAt string                        `json:"updated_at"`
	Comment   string                        `json:"comment"`
	SnippetId string                        `json:"snippet_id"`
	Snippet   resSnip.SnippetCreateResponse `json:"snippet"`
}

func FromUpdateDomain(domain comments.Domain) CommentUpdateResponse {
	t, _ := time.Parse(layoutISO, domain.UpdatedAt.Format(layoutISO))

	return CommentUpdateResponse{
		Id:        domain.Id,
		UpdatedAt: t.Format(layoutUS),
		Username:  domain.Username,
		Comment:   domain.Comment,
		SnippetId: domain.SnippetId,
		Snippet:   resSnip.FromCreateDomain(domain.Snippet),
	}
}
