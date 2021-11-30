package users

import (
	"context"
	"time"
	"twilux/business"
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
		return Domain{}, business.ErrorInvalidEmail
	}
	if domain.Password == "" {
		return Domain{}, business.ErrorInvalidPassword
	}

	user, err := usecase.repo.Login(domain, ctx)

	if !(encrypt.HashValidation(domain.Password, user.Password)) {
		return Domain{}, business.ErrorInvalidPassword
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

// Signup usecase for user
func (usecase *UserUseCase) Register(domain Domain, ctx context.Context) (Domain, error) {
	if domain.Email == "" {
		return Domain{}, business.ErrorInvalidEmail
	}
	if domain.Password == "" {
		return Domain{}, business.ErrorInvalidPassword
	}
	if domain.Username == "" {
		return Domain{}, business.ErrorInvalidUsername
	}
	domain.Password, _ = encrypt.Hash(domain.Password)

	user, err := usecase.repo.Register(domain, ctx)
	if err != nil {
		return Domain{}, err
	}
	return user, nil
}
