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

func NewUsecase(snippetRepo SnippetRepoInterface, contextTimeout time.Duration) SnippetUsecaseInterface {
	return &SnippetUseCase{
		repo: snippetRepo,
		ctx:  contextTimeout,
	}
}

func (usecase *SnippetUseCase) Create(domain Domain, ctx context.Context) (Domain, error) {
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
		return []Domain{}, error
	}
	return snippets, nil
}
