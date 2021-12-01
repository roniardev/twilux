package response

import (
	"twilux/business/snippets"
	format_date "twilux/helpers/date"
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
	t := format_date.FormatDate(domain.UpdatedAt.Format(format_date.LayoutISO))

	return SnippetUpdateResponse{
		Id:        domain.Id,
		UpdatedAt: t,
		Title:     domain.Title,
		Snippet:   domain.Snippet,
		Descb:     domain.Descb,
		Username:  domain.Username,
	}
}
