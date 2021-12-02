package response

import (
	"twilux/business/comments"
	format_date "twilux/helpers/date"
)

type CommentGetResponse struct {
	Id        string `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Username  string `json:"username"`
	Comment   string `json:"comment"`
	SnippetId string `json:"snippet_id"`
}

func FromGetDomain(domain comments.Domain) CommentGetResponse {
	tC := format_date.FormatDate(domain.CreatedAt.Format(format_date.LayoutISO))
	tU := format_date.FormatDate(domain.UpdatedAt.Format(format_date.LayoutISO))

	return CommentGetResponse{
		Id:        domain.Id,
		CreatedAt: tC,
		UpdatedAt: tU,
		Username:  domain.Username,
		Comment:   domain.Comment,
		SnippetId: domain.SnippetId,
	}
}

func ToListDomain(domain []comments.Domain) (response []CommentGetResponse) {
	for _, val := range domain {
		response = append(response, FromGetDomain(val))
	}
	return response
}
