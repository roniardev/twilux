package response

import (
	"time"
	"twilux/business/snippets"
	format_date "twilux/helpers/date"
)

type SnippetDeleteResponse struct {
	Id        string `json:"id"`
	DeletedAt string `json:"deleted_at"`
	Title     string `json:"title"`
	Snippet   string `json:"snippet"`
	Descb     string `json:"descb"`
	Username  string `json:"username"`
}

func FromDeleteDomain(domain snippets.Domain) SnippetDeleteResponse {
	val, _ := domain.DeletedAt.Value()
	t := format_date.FormatDate(val.(time.Time).Format(format_date.LayoutISO))

	return SnippetDeleteResponse{
		Id:        domain.Id,
		DeletedAt: t,
		Title:     domain.Title,
		Snippet:   domain.Snippet,
		Descb:     domain.Descb,
		Username:  domain.Username,
	}
}
