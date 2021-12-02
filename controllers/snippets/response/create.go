package response

import (
	"twilux/business/snippets"
	format_date "twilux/helpers/date"
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
	t := format_date.FormatDate(domain.CreatedAt.Format(format_date.LayoutISO))

	return SnippetCreateResponse{
		Id:        domain.Id,
		CreatedAt: t,
		Title:     domain.Title,
		Snippet:   domain.Snippet,
		Descb:     domain.Descb,
		Username:  domain.Username,
	}
}
