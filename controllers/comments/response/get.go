package response

import (
	"time"
	"twilux/business/comments"
)

type CommentGetResponse struct {
	Id        string `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Username  string `json:"username"`
	Comment   string `json:"comment"`
	SnippetId string `json:"snippet_id"`
}

func FromGetDomain(domain comments.Domain) CommentGetResponse {
	tC, _ := time.Parse(layoutISO, domain.CreatedAt.Format(layoutISO))
	tU, _ := time.Parse(layoutISO, domain.UpdatedAt.Format(layoutISO))

	return CommentGetResponse{
		Id:        domain.Id,
		CreatedAt: tC.Format(layoutUS),
		UpdatedAt: tU.Format(layoutUS),
		Username:  domain.Username,
		Comment:   domain.Comment,
		SnippetId: domain.SnippetId,
	}
}

func ToListDomain(domain []comments.Domain) (response []CommentGetResponse) {
	for _, val := range domain {
		response = append(response, FromGetDomain(val))
	}
	return response
}
