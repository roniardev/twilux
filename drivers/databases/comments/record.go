package comments

import (
	"fmt"
	"time"
	"twilux/business/comments"
	"twilux/drivers/databases/snippets"
	"twilux/drivers/databases/users"

	"gorm.io/gorm"
)

type Comment struct {
	Id        string `gorm:"primaryKey;size:10"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt   `gorm:"index"`
	Comment   string           `gorm:"size:255"`
	SnippetId string           `gorm:"not null;size:10;index"`
	Snippet   snippets.Snippet `gorm:"foreignkey:SnippetId"`
	User      string           `gorm:"not null;size:20;index"`
	UserInfo  users.User       `gorm:"foreignkey:User;references:Username"`
}

func (com Comment) ToDomain() comments.Domain {
	return comments.Domain{
		Id:        com.Id,
		CreatedAt: com.CreatedAt,
		UpdatedAt: com.UpdatedAt,
		DeletedAt: com.DeletedAt,
		Username:  com.User,
		Comment:   com.Comment,
		SnippetId: com.SnippetId,
		Snippet:   com.Snippet.ToDomain(),
	}
}

func FromDomain(domain comments.Domain) Comment {
	return Comment{
		Id:        domain.Id,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
		User:      domain.Username,
		Comment:   domain.Comment,
		SnippetId: domain.SnippetId,
		Snippet:   snippets.FromDomain(domain.Snippet),
	}
}

func ToListDomain(data []Comment) (result []comments.Domain) {
	result = []comments.Domain{}
	fmt.Println(result)
	for _, comment := range data {
		result = append(result, comment.ToDomain())
	}
	return result
}
