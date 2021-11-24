package response

import (
	"fmt"
	"time"
	"twilux/business/comments"
	"twilux/business/snippets"

	"gorm.io/gorm"
)

type CommentResponse struct {
	Id        string          `json:"id"`
	CreatedAt time.Time       `json:"createdAt"`
	UpdatedAt time.Time       `json:"updatedAt"`
	DeletedAt gorm.DeletedAt  `json:"deletedAt"`
	Comment   string          `json:"comment"`
	SnippetId string          `json:"snippetId"`
	Snippet   snippets.Domain `json:"snippet"`
	Username  string          `json:"username"`
}

func FromDomain(domain comments.Domain) CommentResponse {
	return CommentResponse{
		Id:        domain.Id,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
		Comment:   domain.Comment,
		SnippetId: domain.SnippetId,
		Username:  domain.Username,
	}
}

func ToListDomain(domain []comments.Domain) (response []CommentResponse) {
	fmt.Println(domain)
	for _, com := range domain {
		response = append(response, FromDomain(com))
	}
	return response
}
