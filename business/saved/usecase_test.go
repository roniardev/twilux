package saved_test

import (
	"context"
	"os"
	"testing"
	"twilux/business"
	"twilux/business/saved"
	_mockSavedRepository "twilux/business/saved/mocks"
	"twilux/middlewares"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/mock"
)

var (
	savedRepository _mockSavedRepository.SavedRepoInterface
	savedUseCase    saved.SavedUsecaseInterface
	jwtAuth         *middlewares.ConfigJWT
)

func setup() {
	savedUseCase = saved.NewUsecase(&savedRepository, 1)
}

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestSpec(t *testing.T) {
	Convey("Given a saved usecase", t, func() {
		Convey("When create saved", func() {
			saved := saved.Domain{
				Id:        "12dfzc",
				SnippetId: "wsxadae",
			}
			Convey("Succes add saved", func() {
				savedRepository.On("Create", mock.Anything, mock.Anything).Return(saved, nil)
				result, err := savedUseCase.Create(saved, context.Background())
				So(err, ShouldBeNil)
				So(result, ShouldNotBeNil)
				So(result.Id, ShouldEqual, "12dfzc")
				So(result.SnippetId, ShouldEqual, "wsxadae")
			})
			Convey("Failed add saved", func() {
				Convey("empty sbippet_id", func() {
					saved.SnippetId = ""
					savedRepository.On("Create", mock.Anything, mock.Anything).Return(saved, nil)
					_, err := savedUseCase.Create(saved, context.Background())
					So(err, ShouldBeError, business.ErrorInvalidSnippetID)
				})
			})
		})
		// getAll saved
		Convey("When get all saved", func() {
			saved_one := saved.Domain{
				Id:        "wsxadae",
				SnippetId: "asfwx2",
				Username:  "reka",
			}
			saved_two := saved.Domain{
				Id:        "wsxadae",
				SnippetId: "dvd1_f",
				Username:  "reka",
			}
			savedRepository.On("GetAll", mock.Anything, mock.Anything).Return([]saved.Domain{saved_one, saved_two}, nil)
			result, err := savedUseCase.GetAll("reka", context.Background())
			So(err, ShouldBeNil)
			So(result, ShouldNotBeNil)
			So(len(result), ShouldEqual, 2)
		})
		// when delete saved
		Convey("When delete saved", func() {
			saved := saved.Domain{
				SnippetId: "wsxadae",
			}
			Convey("Succes delete", func() {
				savedRepository.On("Delete", mock.Anything, mock.Anything).Return(saved, nil)
				_, err := savedUseCase.Delete(saved, context.Background())
				So(err, ShouldBeNil)
			})
			// error when id is empty
			Convey("Failed delete", func() {
				saved.SnippetId = ""
				savedRepository.On("Delete", mock.Anything, mock.Anything).Return(saved, nil)
				_, err := savedUseCase.Delete(saved, context.Background())
				So(err, ShouldBeError, business.ErrorInvalidSnippetID)
			})
		})
	})
}
