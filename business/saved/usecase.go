package saved

import (
	"context"
	"time"
	"twilux/business"
)

type SavedUseCase struct {
	repo SavedRepoInterface
	ctx  time.Duration
}

func NewUsecase(savedRepo SavedRepoInterface, contextTimeout time.Duration) SavedUsecaseInterface {
	return &SavedUseCase{
		repo: savedRepo,
		ctx:  contextTimeout,
	}
}

func (usecase *SavedUseCase) Create(domain Domain, ctx context.Context) (Domain, error) {
	if domain.SnippetId == "" {
		return Domain{}, business.ErrorInvalidSnippetID
	}
	saved, error := usecase.repo.Create(domain, ctx)

	if error != nil {
		return Domain{}, error
	}
	return saved, nil
}

func (usecase *SavedUseCase) GetAll(username string, ctx context.Context) ([]Domain, error) {
	saveds, error := usecase.repo.GetAll(username, ctx)

	if error != nil {
		return []Domain{}, error
	}
	return saveds, nil
}

// Delete Saved
func (usecase *SavedUseCase) Delete(domain Domain, ctx context.Context) (Domain, error) {
	if domain.SnippetId == "" {
		return Domain{}, business.ErrorInvalidSnippetID
	}

	saved, error := usecase.repo.Delete(domain, ctx)

	if error != nil {
		return Domain{}, error
	}
	return saved, nil
}
