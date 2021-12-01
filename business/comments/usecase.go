package comments

import (
	"context"
	"time"
	"twilux/business"
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
		return Domain{}, business.ErrorInvalidSnippetID
	}
	if domain.Comment == "" {
		return Domain{}, business.ErrorEmptyComment
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

func (usecase *CommentUseCase) GetAllUser(username string, ctx context.Context) ([]Domain, error) {
	saveds, error := usecase.repo.GetAllUser(username, ctx)

	if error != nil {
		return []Domain{}, error
	}
	return saveds, nil
}

func (usecase *CommentUseCase) Update(domain Domain, ctx context.Context) (Domain, error) {
	if domain.Id == "" {
		return Domain{}, business.ErrorInvalidCommentID
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
		return Domain{}, business.ErrorInvalidCommentID
	}

	saved, error := usecase.repo.Delete(domain, ctx)

	if error != nil {
		return Domain{}, error
	}
	return saved, nil
}
