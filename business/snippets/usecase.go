package snippets

import (
	"context"
	"time"
	"twilux/business"
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
		return Domain{}, business.ErrorEmptyTitle
	}
	if domain.Snippet == "" {
		return Domain{}, business.ErrorEmptySnippet
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
		return Domain{}, business.ErrorInvalidSnippetID
	}

	snippet, error := usecase.repo.Delete(domain, ctx)

	if error != nil {
		return Domain{}, error
	}
	return snippet, nil
}
