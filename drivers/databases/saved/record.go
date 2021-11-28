package saved

import (
	"fmt"
	"time"
	"twilux/business/saved"
	"twilux/drivers/databases/snippets"

	"twilux/drivers/databases/users"

	"gorm.io/gorm"
)

type Saved struct {
	Id        string `gorm:"primaryKey;size:10"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt   `gorm:"index"`
	SnippetId string           `gorm:"not null;size:10;index"`
	Snippet   snippets.Snippet `gorm:"foreignkey:SnippetId"`
	User      string           `gorm:"not null;size:20;index"`
	UserInfo  users.User       `gorm:"foreignkey:User;references:Username"`
}

func (sav Saved) ToDomain() saved.Domain {
	return saved.Domain{
		Id:        sav.Id,
		CreatedAt: sav.CreatedAt,
		UpdatedAt: sav.UpdatedAt,
		DeletedAt: sav.DeletedAt,
		SnippetId: sav.SnippetId,
		Snippet:   sav.Snippet.ToDomain(),
		Username:  sav.User,
	}
}

func FromDomain(domain saved.Domain) Saved {
	return Saved{
		Id:        domain.Id,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
		SnippetId: domain.SnippetId,
		Snippet:   snippets.FromDomain(domain.Snippet),
		User:      domain.Username,
	}
}

func ToListDomain(data []Saved) (result []saved.Domain) {
	result = []saved.Domain{}
	fmt.Println(result)
	for _, sav := range data {
		result = append(result, sav.ToDomain())
	}
	return result
}
