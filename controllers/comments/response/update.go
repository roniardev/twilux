package response

import (
	"twilux/business/comments"
	resSnip "twilux/controllers/snippets/response"
	format_date "twilux/helpers/date"
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
	t := format_date.FormatDate(domain.UpdatedAt.Format(format_date.LayoutISO))

	return CommentUpdateResponse{
		Id:        domain.Id,
		UpdatedAt: t,
		Username:  domain.Username,
		Comment:   domain.Comment,
		SnippetId: domain.SnippetId,
		Snippet:   resSnip.FromCreateDomain(domain.Snippet),
	}
}
