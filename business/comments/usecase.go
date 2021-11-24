package comments

import (
	"context"
	"errors"
	"time"
)

type CommentUseCase struct {
	repo CommentRepoInterface
	ctx  time.Duration
}

func NewUsecase(savedRepo CommentRepoInterface, contextTimeout time.Duration) CommentUsecaseInterface {
	return &CommentUseCase{
		repo: savedRepo,
		ctx:  contextTimeout,
	}
}

func (usecase *CommentUseCase) Create(domain Domain, ctx context.Context) (Domain, error) {
	if domain.SnippetId == "" {
		return Domain{}, errors.New("snippet id is required")
	}
	if domain.Username == "" {
		return Domain{}, errors.New("username is required")
	}
	saved, error := usecase.repo.Create(domain, ctx)

	if error != nil {
		return Domain{}, error
	}
	return saved, nil
}

func (usecase *CommentUseCase) GetAll(snippetId string, ctx context.Context) ([]Domain, error) {
	saveds, error := usecase.repo.GetAll(snippetId, ctx)

	if error != nil {
		return []Domain{}, error
	}
	return saveds, nil
}

func (usecase *CommentUseCase) Update(domain Domain, ctx context.Context) (Domain, error) {
	if domain.Id == "" {
		return Domain{}, errors.New("id is required")
	}
	if domain.Comment == "" {
		return Domain{}, errors.New("comment is required")
	}
	snippet, error := usecase.repo.Update(domain, ctx)

	if error != nil {
		return Domain{}, error
	}
	return snippet, nil
}

// Delete Comment
func (usecase *CommentUseCase) Delete(domain Domain, ctx context.Context) (Domain, error) {
	if domain.Id == "" {
		return Domain{}, errors.New("id is required")
	}

	saved, error := usecase.repo.Delete(domain, ctx)

	if error != nil {
		return Domain{}, error
	}
	return saved, nil
}
