package response

import (
	"fmt"
	"time"
	"twilux/business/saved"
	resSnip "twilux/controllers/snippets/response"
)

const (
	layoutISO = "2006-01-02T15:04:05.999999999Z07:00"
	layoutUS  = "Monday, January 2, 2006, 03:04:01 PM"
)

type SavedResponse struct {
	Id        string                        `json:"id"`
	CreatedAt string                        `json:"created_at"`
	SnippetId string                        `json:"snippet_id"`
	Snippet   resSnip.SnippetCreateResponse `json:"snippet"`
	Username  string                        `json:"username"`
}

func FromDomain(domain saved.Domain) SavedResponse {
	t, _ := time.Parse(layoutISO, domain.CreatedAt.Format(layoutISO))

	return SavedResponse{
		CreatedAt: t.Format(layoutUS),
		Id:        domain.Id,
		SnippetId: domain.SnippetId,
		Snippet:   resSnip.FromCreateDomain(domain.Snippet),
		Username:  domain.Username,
	}
}

func ToListDomain(domain []saved.Domain) (response []SavedResponse) {
	fmt.Println(domain)
	for _, sav := range domain {
		response = append(response, FromDomain(sav))
	}
	return response
}
