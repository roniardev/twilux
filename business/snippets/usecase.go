package snippets

import (
	"context"
	"errors"
	"time"
)

type SnippetUseCase struct {
	repo SnippetRepoInterface
	ctx  time.Duration
}

func NewUsecase(snippetRepo SnippetRepoInterface, ctx time.Duration) *SnippetUseCase {
	return &SnippetUseCase{
		repo: snippetRepo,
		ctx:  ctx,
	}
}

func (usecase *SnippetUseCase) Create(ctx context.Context, domain Domain) (Domain, error) {
	if domain.Title == "" {
		return Domain{}, errors.New("title is required")
	}
	if domain.Snippet == "" {
		return Domain{}, errors.New("snippet is required")
	}
	snippet, error := usecase.repo.Create(domain, ctx)

	if error != nil {
		return Domain{}, error
	}
	return snippet, nil
}

func (usecase *SnippetUseCase) GetAll(ctx context.Context) ([]Domain, error) {
	snippets, error := usecase.repo.GetAll(ctx)

	if error != nil {
		return nil, error
	}
	return snippets, nil
}
