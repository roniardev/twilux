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

func (usecase *SnippetUseCase) GetById(snippetId string, ctx context.Context) (Domain, error) {
	snippet, error := usecase.repo.GetById(snippetId, ctx)

	if error != nil {
		return Domain{}, error
	}
	return snippet, nil
}

// Update func updates a snippet
func (usecase *SnippetUseCase) Update(domain Domain, ctx context.Context) (Domain, error) {
	snippet, error := usecase.repo.Update(domain, ctx)

	if error != nil {
		return Domain{}, error
	}
	return snippet, nil
}

// Delete Snippet
func (usecase *SnippetUseCase) Delete(domain Domain, ctx context.Context) (Domain, error) {
	if domain.Id == "" {
		return Domain{}, errors.New("id is required")
	}

	snippet, error := usecase.repo.Delete(domain, ctx)

	if error != nil {
		return Domain{}, error
	}
	return snippet, nil
}
