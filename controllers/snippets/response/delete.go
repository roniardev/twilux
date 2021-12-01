package response

import (
	"time"
	"twilux/business/snippets"
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
	t, _ := time.Parse(layoutISO, val.(time.Time).Format(layoutISO))

	return SnippetDeleteResponse{
		Id:        domain.Id,
		DeletedAt: t.Format(layoutUS),
		Title:     domain.Title,
		Snippet:   domain.Snippet,
		Descb:     domain.Descb,
		Username:  domain.Username,
	}
}
