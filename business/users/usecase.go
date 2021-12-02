package users

import (
	"context"
	"strings"
	"time"
	"twilux/business"
	"twilux/helpers/encrypt"
	"twilux/middlewares"

	emailverifier "github.com/AfterShip/email-verifier"
)

var (
	verifier = emailverifier.NewVerifier().EnableAutoUpdateDisposable()
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
		return Domain{}, business.ErrorEmptyEmail
	}
	if domain.Password == "" {
		return Domain{}, business.ErrorEmptyPassword
	}

	user, err := usecase.repo.Login(domain, ctx)
	if err != nil {
		return Domain{}, business.ErrorDataNotFound
	}

	if !(encrypt.HashValidation(domain.Password, user.Password)) {
		return Domain{}, business.ErrorInvalidPassword
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
		return Domain{}, business.ErrorEmptyEmail
	}
	if domain.Password == "" {
		return Domain{}, business.ErrorEmptyPassword
	}
	if domain.Username == "" {
		return Domain{}, business.ErrorInvalidUsername
	}

	emailDomain := strings.Split(domain.Email, "@")[1]

	_, verifEmail := verifier.CheckMX(emailDomain)
	isDispoEmail := verifier.IsDisposable(emailDomain)
	if verifEmail != nil || isDispoEmail {
		return Domain{}, business.ErrorInvalidEmail
	}

	domain.Password, _ = encrypt.Hash(domain.Password)

	user, err := usecase.repo.Register(domain, ctx)
	if err != nil {
		return Domain{}, err
	}
	return user, nil
}
