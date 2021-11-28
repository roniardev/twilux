package response

import (
	"time"
	"twilux/business/saved"
)

type SavedDeleteResponse struct {
	Id        string `json:"id"`
	DeletedAt string `json:"deleted_at"`
	SnippetId string `json:"snippet_id"`
	Username  string `json:"username"`
}

func FromDeleteDomain(domain saved.Domain) SavedDeleteResponse {
	val, _ := domain.DeletedAt.Value()
	t, _ := time.Parse(layoutISO, val.(time.Time).Format(layoutISO))

	return SavedDeleteResponse{
		Id:        domain.Id,
		DeletedAt: t.Format(layoutUS),
		SnippetId: domain.SnippetId,
		Username:  domain.Username,
	}
}
