package response

import (
	"time"
	"twilux/business/snippets"
)

type SnippetUpdateResponse struct {
	Id        string `json:"id"`
	UpdatedAt string `json:"created_at"`
	Title     string `json:"title"`
	Snippet   string `json:"snippet"`
	Descb     string `json:"descb"`
	Username  string `json:"username"`
}

func FromUpdateDomain(domain snippets.Domain) SnippetUpdateResponse {
	t, _ := time.Parse(layoutISO, domain.UpdatedAt.Format(layoutISO))

	return SnippetUpdateResponse{
		Id:        domain.Id,
		UpdatedAt: t.Format(layoutUS),
		Title:     domain.Title,
		Snippet:   domain.Snippet,
		Descb:     domain.Descb,
		Username:  domain.Username,
	}
}
