package response

import (
	"time"
	"twilux/business/snippets"
)

const (
	layoutISO = "2006-01-02T15:04:05.999999999Z07:00"
	layoutUS  = "Monday, January 2, 2006, 03:04:01 PM"
)

type SnippetCreateResponse struct {
	Id        string `json:"id"`
	CreatedAt string `json:"created_at"`
	Title     string `json:"title"`
	Snippet   string `json:"snippet"`
	Descb     string `json:"descb"`
	Username  string `json:"username"`
}

func FromCreateDomain(domain snippets.Domain) SnippetCreateResponse {
	t, _ := time.Parse(layoutISO, domain.CreatedAt.Format(layoutISO))

	return SnippetCreateResponse{
		Id:        domain.Id,
		CreatedAt: t.Format(layoutUS),
		Title:     domain.Title,
		Snippet:   domain.Snippet,
		Descb:     domain.Descb,
		Username:  domain.Username,
	}
}
