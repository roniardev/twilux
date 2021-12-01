package snippets_test

import (
	"context"
	"os"
	"testing"
	"twilux/business"
	"twilux/business/snippets"
	_mockSnippetRepository "twilux/business/snippets/mocks"
	"twilux/middlewares"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/mock"
)

var (
	snippetRepository _mockSnippetRepository.SnippetRepoInterface
	snippetUseCase    snippets.SnippetUsecaseInterface
	jwtAuth           *middlewares.ConfigJWT
)

func setup() {
	snippetUseCase = snippets.NewUsecase(&snippetRepository, 1)
}

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestSpec(t *testing.T) {
	Convey("Given a snippet usecase", t, func() {
		Convey("When create snippet", func() {
			snippet := snippets.Domain{
				Id:      "wsxadae",
				Title:   "title",
				Snippet: "content",
				Descb:   "go",
			}
			Convey("Succes add snippet", func() {
				snippetRepository.On("Create", mock.Anything, mock.Anything).Return(snippet, nil)
				result, err := snippetUseCase.Create(snippet, context.Background())
				So(err, ShouldBeNil)
				So(result, ShouldNotBeNil)
				So(result.Id, ShouldEqual, "wsxadae")
				So(result.Title, ShouldEqual, "title")
				So(result.Snippet, ShouldEqual, "content")
				So(result.Descb, ShouldEqual, "go")
			})
			Convey("Failed add snippet", func() {
				Convey("empty title", func() {
					snippet.Title = ""
					snippetRepository.On("Create", mock.Anything, mock.Anything).Return(snippets.Domain{}, nil)
					_, err := snippetUseCase.Create(snippet, context.Background())
					So(err, ShouldBeError, business.ErrorEmptyTitle)
				})
				Convey("empty snippet", func() {
					snippet.Snippet = ""
					snippetRepository.On("Create", mock.Anything, mock.Anything).Return(snippets.Domain{}, nil)
					_, err := snippetUseCase.Create(snippet, context.Background())
					So(err, ShouldBeError, business.ErrorEmptySnippet)
				})
			})
		})
		// getAll snippet
		Convey("When get all snippet", func() {
			snippet_one := snippets.Domain{
				Id:      "wsxadae",
				Title:   "title",
				Snippet: "content",
				Descb:   "go",
			}
			snippet_two := snippets.Domain{
				Id:      "wsxadae",
				Title:   "title",
				Snippet: "content",
				Descb:   "go",
			}
			snippetRepository.On("GetAll", mock.Anything, mock.Anything).Return([]snippets.Domain{snippet_one, snippet_two}, nil)
			result, err := snippetUseCase.GetAll(context.Background())
			So(err, ShouldBeNil)
			So(result, ShouldNotBeNil)
			So(len(result), ShouldEqual, 2)
		})
		// get snippet by id
		Convey("When get snippet by id", func() {
			snippet := snippets.Domain{
				Id:      "wsxadae",
				Title:   "title",
				Snippet: "content",
				Descb:   "go",
			}
			snippetRepository.On("GetById", mock.Anything, mock.Anything).Return(snippet, nil)
			result, err := snippetUseCase.GetById("wsxadae", context.Background())
			So(err, ShouldBeNil)
			So(result, ShouldResemble, snippet)
			So(result.Id, ShouldEqual, snippet.Id)
		})
		// update snippet
		Convey("When update snippet", func() {
			snippet := snippets.Domain{
				Id:      "wsxadae",
				Title:   "title",
				Snippet: "content",
				Descb:   "go",
			}
			// success update
			Convey("Succes update", func() {
				updated_snippet := &snippet
				updated_snippet.Title = "new title"
				snippetRepository.On("Update", mock.Anything, mock.Anything).Return(snippet, nil)
				result, err := snippetUseCase.Update(*updated_snippet, context.Background())
				So(err, ShouldBeNil)
				So(result, ShouldNotBeNil)
				So(result.Id, ShouldEqual, "wsxadae")
				So(snippet.Title, ShouldEqual, "new title")
			})
		})
		// when delete snippet
		Convey("When delete snippet", func() {
			snippet := snippets.Domain{
				Id: "wsxadae",
			}
			Convey("Succes delete", func() {
				snippetRepository.On("Delete", mock.Anything, mock.Anything).Return(snippet, nil)
				_, err := snippetUseCase.Delete(snippet, context.Background())
				So(err, ShouldBeNil)
			})
			// error when id is empty
			Convey("Failed delete", func() {
				snippet.Id = ""
				snippetRepository.On("Delete", mock.Anything, mock.Anything).Return(snippet, nil)
				_, err := snippetUseCase.Delete(snippet, context.Background())
				So(err, ShouldBeError, business.ErrorInvalidSnippetID)
			})
		})
	})
}
