package response

import (
	"fmt"
	"twilux/business/saved"
	resSnip "twilux/controllers/snippets/response"
	format_date "twilux/helpers/date"
)

type SavedResponse struct {
	Id        string                        `json:"id"`
	CreatedAt string                        `json:"created_at"`
	SnippetId string                        `json:"snippet_id"`
	Snippet   resSnip.SnippetCreateResponse `json:"snippet"`
	Username  string                        `json:"username"`
}

func FromDomain(domain saved.Domain) SavedResponse {
	t := format_date.FormatDate(domain.CreatedAt.Format(format_date.LayoutISO))

	return SavedResponse{
		CreatedAt: t,
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
