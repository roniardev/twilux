package users

import (
	"context"
	"errors"
	"fmt"
	"time"
	"twilux/helpers/encrypt"
	"twilux/middlewares"
)

type UserUseCase struct {
	ConfigJWT middlewares.ConfigJWT
	repo      UserRepoInterface
	ctx       time.Duration
}

func NewUsecase(configJWT middlewares.ConfigJWT, userRepo UserRepoInterface, contextTimeout time.Duration) UserUsecaseInterface {
	return &UserUseCase{
		ConfigJWT: configJWT,
		repo:      userRepo,
		ctx:       contextTimeout,
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

	if !(encrypt.HashValidation(domain.Password, user.Password)) {
		return Domain{}, errors.New("invalid password")
	}

	if err != nil {
		return Domain{}, err
	}

	user.Token, err = usecase.ConfigJWT.GenerateToken(user.Username)
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
		return Domain{}, errors.New("email required")
	}
	if domain.Password == "" {
		return Domain{}, errors.New("password required")
	}
	if domain.Username == "" {
		return Domain{}, errors.New("username required")
	}
	domain.Password, _ = encrypt.Hash(domain.Password)

	user, err := usecase.repo.Register(domain, ctx)
	if err != nil {
		return Domain{}, err
	}
	return user, nil
}
