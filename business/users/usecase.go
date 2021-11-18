package users

import (
	"context"
	"errors"
	"fmt"
	"time"
)

type UserUseCase struct {
	repo UserRepoInterface
	ctx  time.Duration
}

func NewUsecase(userRepo UserRepoInterface, contextTimeout time.Duration) UserUsecaseInterface {
	return &UserUseCase{
		repo: userRepo,
		ctx:  contextTimeout,
	}
}

func (usecase *UserUseCase) Login(domain Domain, ctx context.Context) (Domain, error) {
	if domain.Email == "" {
		fmt.Println("email Empty")
		return Domain{}, errors.New("email required")
	}
	if domain.Password == "" {
		return Domain{}, errors.New("password required")
	}
	user, err := usecase.repo.Login(domain, ctx)
	if err != nil {
		return Domain{}, err
	}
	return user, nil
}

func (usecase *UserUseCase) GetAllUsers(ctx context.Context) ([]Domain, error) {
	return []Domain{}, nil
}

// Signup usecase for user
func (usecase *UserUseCase) Register(domain Domain, ctx context.Context) (Domain, error) {
	if domain.Email == "" {
		return Domain{}, errors.New("email Empty")
	}
	if domain.Password == "" {
		return Domain{}, errors.New("password Empty")
	}
	if domain.Username == "" {
		return Domain{}, errors.New("name Empty")
	}

	user, err := usecase.repo.Register(domain, ctx)
	if err != nil {
		return Domain{}, err
	}
	return user, nil
}
