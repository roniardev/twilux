package response

import (
	"time"
	"twilux/business/comments"
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
	t, _ := time.Parse(layoutISO, val.(time.Time).Format(layoutISO))

	return CommentDeleteResponse{
		Id:        domain.Id,
		DeletedAt: t.Format(layoutUS),
		Username:  domain.Username,
		Comment:   domain.Comment,
		SnippetId: domain.SnippetId,
	}
}
