package response

import (
	"twilux/business/comments"
	resSnip "twilux/controllers/snippets/response"
	format_date "twilux/helpers/date"
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
	t := format_date.FormatDate(domain.CreatedAt.Format(format_date.LayoutISO))

	return CommentResponse{
		CreatedAt: t,
		Id:        domain.Id,
		Username:  domain.Username,
		Comment:   domain.Comment,
		SnippetId: domain.SnippetId,
		Snippet:   resSnip.FromCreateDomain(domain.Snippet),
	}
}
