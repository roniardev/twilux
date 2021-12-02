package response

import (
	"time"
	"twilux/business/comments"
	format_date "twilux/helpers/date"
)

type CommentDeleteResponse struct {
	Id        string `json:"id"`
	DeletedAt string `json:"deleted_at"`
	Username  string `json:"username"`
	Comment   string `json:"comment"`
	SnippetId string `json:"snippet_id"`
}

func FromDeleteDomain(domain comments.Domain) CommentDeleteResponse {
	val, _ := domain.DeletedAt.Value()
	t := format_date.FormatDate(val.(time.Time).Format(format_date.LayoutISO))

	return CommentDeleteResponse{
		Id:        domain.Id,
		DeletedAt: t,
		Username:  domain.Username,
		Comment:   domain.Comment,
		SnippetId: domain.SnippetId,
	}
}
