package users_test

import (
	"context"
	"os"
	"testing"
	"twilux/business"
	"twilux/business/users"
	_mockUserRepository "twilux/business/users/mocks"
	"twilux/helpers/encrypt"
	"twilux/middlewares"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/mock"
)

var (
	userRepository _mockUserRepository.UserRepoInterface
	userUseCase    users.UserUsecaseInterface
	jwtAuth        *middlewares.ConfigJWT
)

func setup() {
	jwtAuth = &middlewares.ConfigJWT{SecretJWT: "twistedFateLux", ExpiresDuration: 1}
	userUseCase = users.NewUsecase(*jwtAuth, &userRepository, 1)
}

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestSpec(t *testing.T) {
	// Convey register usecase test
	Convey("Given a user repository", t, func() {
		Convey("When register a user", func() {
			hashedPassword, _ := encrypt.Hash("112233")
			domainUser := users.Domain{
				Email:    "example@gmail.com",
				Password: "test",
				Username: hashedPassword,
			}
			userRepository.On("Register", mock.Anything, mock.Anything).Return(domainUser, nil).Once()
			val, err := userUseCase.Register(domainUser, context.Background())
			Convey("Should not return error", func() {
				So(err, ShouldBeNil)
			})
			Convey("Should return user", func() {
				So(val, ShouldResemble, domainUser)
			})
			Convey("Error when user request with empty value", func() {
				Convey("Error when email invalid or empty", func() {
					domainUser.Email = ""
					_, err := userUseCase.Register(domainUser, context.Background())
					So(err, ShouldBeError, business.ErrorEmptyEmail)
				})
				// error when invalid username
				Convey("Error when username invalid or empty", func() {
					domainUser.Username = ""
					_, err := userUseCase.Register(domainUser, context.Background())
					So(err, ShouldBeError, business.ErrorInvalidUsername)
				})
				// error when invalid password
				Convey("Error when password invalid or empty", func() {
					domainUser.Password = ""
					_, err := userUseCase.Register(domainUser, context.Background())
					So(err, ShouldBeError, business.ErrorEmptyPassword)
				})
				Convey("Error when user using invalid email ex: spam or disposal email", func() {
					domainUser.Email = "reka@tempmail.eu"
					_, err := userUseCase.Register(domainUser, context.Background())
					So(err, ShouldBeError, business.ErrorInvalidEmail)
				})
			})

		})
		// when login user
		Convey("When login a user", func() {
			hashedPassword, _ := encrypt.Hash("112233")

			domainUser := users.Domain{
				Email:    "test@test.com",
				Password: hashedPassword,
			}
			userRepository.On("Login", mock.Anything, mock.Anything).Return(domainUser, nil).Once()
			Convey("Login success", func() {
				_, err := userUseCase.Login(users.Domain{
					Email:    "test@test.com",
					Password: "112233",
				}, context.Background())
				So(err, ShouldBeNil)
			})

			Convey("Error when email empty", func() {
				domainUser.Email = ""
				_, err := userUseCase.Login(domainUser, context.Background())
				So(err, ShouldBeError, business.ErrorEmptyEmail)
			})
			Convey("Error when password empty", func() {
				domainUser.Password = ""
				_, err := userUseCase.Login(domainUser, context.Background())
				So(err, ShouldBeError, business.ErrorEmptyPassword)
			})
			Convey("Error when invalid password", func() {
				_, err := userUseCase.Login(users.Domain{
					Email:    "test@test.com",
					Password: "leh",
				}, context.Background())
				So(err, ShouldBeError, business.ErrorInvalidPassword)
			})
		})
	})
}
