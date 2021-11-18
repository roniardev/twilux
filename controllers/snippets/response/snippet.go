package response

import (
	"time"
	"twilux/business/snippets"

	"gorm.io/gorm"
)

type SnippetResponse struct {
	Id        string         `json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
	Title     string         `json:"title"`
	Snippet   string         `json:"snippet"`
	Descb     string         `json:"descb"`
	Username  string         `json:"username"`
}

func FromDomain(domain snippets.Domain) SnippetResponse {
	return SnippetResponse{
		Id:        domain.Id,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
		Title:     domain.Title,
		Snippet:   domain.Snippet,
		Descb:     domain.Descb,
		Username:  domain.Username,
	}
}
