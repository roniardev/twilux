package response

import (
	"fmt"
	"time"
	"twilux/business/snippets"
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
	tC, _ := time.Parse(layoutISO, domain.CreatedAt.Format(layoutISO))
	tU, _ := time.Parse(layoutISO, domain.UpdatedAt.Format(layoutISO))

	return SnippetResponse{
		Id:        domain.Id,
		CreatedAt: tC.Format(layoutUS),
		UpdatedAt: tU.Format(layoutUS),
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
