package response

import (
	"fmt"
	"time"
	"twilux/business/saved"
	"twilux/business/snippets"

	"gorm.io/gorm"
)

type SavedResponse struct {
	Id        string            `json:"id"`
	CreatedAt time.Time         `json:"createdAt"`
	UpdatedAt time.Time         `json:"updatedAt"`
	DeletedAt gorm.DeletedAt    `json:"deletedAt"`
	SnippetId string            `json:"snippetId"`
	Snippet   []snippets.Domain `json:"snippet"`
	Username  string            `json:"username"`
}

func FromDomain(domain saved.Domain) SavedResponse {
	return SavedResponse{
		Id:        domain.Id,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
		SnippetId: domain.SnippetId,
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
