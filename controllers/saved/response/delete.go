package response

import (
	"time"
	"twilux/business/saved"
	format_date "twilux/helpers/date"
)

type SavedDeleteResponse struct {
	Id        string `json:"id"`
	DeletedAt string `json:"deleted_at"`
	SnippetId string `json:"snippet_id"`
	Username  string `json:"username"`
}

func FromDeleteDomain(domain saved.Domain) SavedDeleteResponse {
	val, _ := domain.DeletedAt.Value()
	t := format_date.FormatDate(val.(time.Time).Format(format_date.LayoutISO))

	return SavedDeleteResponse{
		Id:        domain.Id,
		DeletedAt: t,
		SnippetId: domain.SnippetId,
		Username:  domain.Username,
	}
}
