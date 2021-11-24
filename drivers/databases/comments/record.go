package comments

import (
	"fmt"
	"time"
	"twilux/business/comments"

	"gorm.io/gorm"
)

type Comment struct {
	Id        string `gorm:"primaryKey;size:10"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Comment   string         `gorm:"size:255"`
	SnippetId string         `gorm:"not null;size:10;index"`
	Username  string         `gorm:"not null;size:20;index"`
}

func (com Comment) ToDomain() comments.Domain {
	return comments.Domain{
		Id:        com.Id,
		CreatedAt: com.CreatedAt,
		UpdatedAt: com.UpdatedAt,
		DeletedAt: com.DeletedAt,
		Comment:   com.Comment,
		SnippetId: com.SnippetId,
		Username:  com.Username,
	}
}

func FromDomain(domain comments.Domain) Comment {
	return Comment{
		Id:        domain.Id,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
		Comment:   domain.Comment,
		SnippetId: domain.SnippetId,
		Username:  domain.Username,
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
