package response

import (
	"fmt"
	"twilux/business/snippets"
	format_date "twilux/helpers/date"
)

type SnippetResponse struct {
	Id        string `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Title     string `json:"title"`
	Snippet   string `json:"snippet"`
	Descb     string `json:"descb"`
	Username  string `json:"username"`
}

func FromDomain(domain snippets.Domain) SnippetResponse {
	tC := format_date.FormatDate(domain.CreatedAt.Format(format_date.LayoutISO))
	tU := format_date.FormatDate(domain.UpdatedAt.Format(format_date.LayoutISO))

	return SnippetResponse{
		Id:        domain.Id,
		CreatedAt: tC,
		UpdatedAt: tU,
		Title:     domain.Title,
		Snippet:   domain.Snippet,
		Descb:     domain.Descb,
		Username:  domain.Username,
	}
}

func ToListDomain(domain []snippets.Domain) (response []SnippetResponse) {
	fmt.Println(domain)
	for _, snippet := range domain {
		response = append(response, FromDomain(snippet))
	}
	return response
}
