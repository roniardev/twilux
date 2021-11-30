package comments_test

import (
	"context"
	"os"
	"testing"
	"twilux/business"
	"twilux/business/comments"
	_mockCommentRepository "twilux/business/comments/mocks"
	"twilux/middlewares"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/mock"
)

var (
	commentRepository _mockCommentRepository.CommentRepoInterface
	commentUseCase    comments.CommentUsecaseInterface
	jwtAuth           *middlewares.ConfigJWT
)

func setup() {
	commentUseCase = comments.NewUsecase(&commentRepository, 1)
}

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestSpec(t *testing.T) {
	Convey("Given a comment usecase", t, func() {
		Convey("When create comment", func() {
			comment := comments.Domain{
				Id:        "wsxadae",
				SnippetId: "adau7",
				Comment:   "content",
			}
			Convey("Succes add comment", func() {
				commentRepository.On("Create", mock.Anything, mock.Anything).Return(comment, nil)
				result, err := commentUseCase.Create(comment, context.Background())
				So(err, ShouldBeNil)
				So(result, ShouldNotBeNil)
				So(result.Id, ShouldEqual, "wsxadae")
				So(result.SnippetId, ShouldEqual, "adau7")
				So(result.Comment, ShouldEqual, "content")
			})
			Convey("Failed add comment", func() {
				Convey("empty title", func() {
					comment.SnippetId = ""
					commentRepository.On("Create", mock.Anything, mock.Anything).Return(comments.Domain{}, nil)
					_, err := commentUseCase.Create(comment, context.Background())
					So(err, ShouldBeError, business.ErrorInvalidSnippetID)
				})
				Convey("empty comment", func() {
					comment.Comment = ""
					commentRepository.On("Create", mock.Anything, mock.Anything).Return(comments.Domain{}, nil)
					_, err := commentUseCase.Create(comment, context.Background())
					So(err, ShouldBeError, business.ErrorEmptyComment)
				})
			})
		})
		// getAll comment
		Convey("When get all comment", func() {
			comment_one := comments.Domain{
				Id:        "wsxadae",
				SnippetId: "title",
				Comment:   "darius",
			}
			comment_two := comments.Domain{
				Id:        "wsxadae",
				SnippetId: "title",
				Comment:   "content",
				Username:  "darius",
			}
			comment_three := comments.Domain{
				Id:        "wsxadae",
				SnippetId: "title",
				Comment:   "content",
				Username:  "garen",
			}
			commentRepository.On("GetAll", mock.Anything, mock.Anything).Return([]comments.Domain{comment_one, comment_two, comment_three}, nil)
			result, err := commentUseCase.GetAll("wsxadae", context.Background())
			So(err, ShouldBeNil)
			So(result, ShouldNotBeNil)
			So(len(result), ShouldEqual, 3)
		})
		// get comment by id
		Convey("When get all comment by username", func() {
			comment_one := comments.Domain{
				Id:        "wsxadae",
				SnippetId: "title",
				Comment:   "darius",
			}
			comment_two := comments.Domain{
				Id:        "wsxadae",
				SnippetId: "title",
				Comment:   "content",
				Username:  "darius",
			}
			comment_three := comments.Domain{
				Id:        "wsxadae",
				SnippetId: "title",
				Comment:   "content",
				Username:  "garen",
			}
			commentRepository.On("GetAllUser", mock.Anything, mock.Anything).Return([]comments.Domain{comment_one, comment_two, comment_three}, nil)
			result, err := commentUseCase.GetAllUser("wsxadae", context.Background())
			So(err, ShouldBeNil)
			So(result, ShouldNotBeNil)
			So(len(result), ShouldEqual, 3)
		})
		// update comment
		Convey("When update comment", func() {
			comment := comments.Domain{
				Id:        "wsxadae",
				SnippetId: "fsfsf5f",
				Comment:   "Heiho good code",
			}
			// success update
			Convey("Succes update", func() {
				updated_comment := &comment
				updated_comment.Comment = "thx bro for the code"
				commentRepository.On("Update", mock.Anything, mock.Anything).Return(comment, nil)
				result, err := commentUseCase.Update(*updated_comment, context.Background())
				So(err, ShouldBeNil)
				So(result, ShouldNotBeNil)
				So(result.Id, ShouldEqual, "wsxadae")
				So(result.SnippetId, ShouldEqual, "fsfsf5f")
				So(comment.Comment, ShouldEqual, "thx bro for the code")
			})
			Convey("Failed update", func() {
				comment.Id = ""
				updated_comment := &comment
				updated_comment.Comment = "thx bro for the code"
				commentRepository.On("Delete", mock.Anything, mock.Anything).Return(comment, nil)
				_, err := commentUseCase.Update(comment, context.Background())
				So(err, ShouldBeError, business.ErrorInvalidCommentID)
			})
		})
		// when delete comment
		Convey("When delete comment", func() {
			comment := comments.Domain{
				Id: "wsxadae",
			}
			Convey("Succes delete", func() {
				commentRepository.On("Delete", mock.Anything, mock.Anything).Return(comment, nil)
				_, err := commentUseCase.Delete(comment, context.Background())
				So(err, ShouldBeNil)
			})
			// error when id is empty
			Convey("Failed delete", func() {
				comment.Id = ""
				commentRepository.On("Delete", mock.Anything, mock.Anything).Return(comment, nil)
				_, err := commentUseCase.Delete(comment, context.Background())
				So(err, ShouldBeError, business.ErrorInvalidCommentID)
			})
		})
	})
}
