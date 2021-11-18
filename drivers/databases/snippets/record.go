package snippets

import (
	"time"
	"twilux/business/snippets"

	"github.com/jkomyno/nanoid"
	"gorm.io/gorm"
)

type Snippet struct {
	Id        string `gorm:"primaryKey;size:10"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Title     string         `gorm:"size:50;not null"`
	Descb     string
	Snippet   string `gorm:"not null"`
	Username  string `gorm:"not null;size:20;index"`
}

func (snippet Snippet) ToDomain() snippets.Domain {
	return snippets.Domain{
		Id:        snippet.Id,
		CreatedAt: snippet.CreatedAt,
		UpdatedAt: snippet.UpdatedAt,
		DeletedAt: snippet.DeletedAt,
		Title:     snippet.Title,
		Descb:     snippet.Descb,
		Snippet:   snippet.Snippet,
		Username:  snippet.Username,
	}
}

func FromDomain(domain snippets.Domain) Snippet {
	val, _ := nanoid.Nanoid(10)
	return Snippet{
		Id:        val,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
		Title:     domain.Title,
		Descb:     domain.Descb,
		Snippet:   domain.Snippet,
		Username:  domain.Username,
	}
}
