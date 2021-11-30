package business

import "errors"

var (
	ErrorInvalidPassword  = errors.New("password cannot be null or empty")
	ErrorInvalidEmail     = errors.New("email cannot be null or empty")
	ErrorInvalidUsername  = errors.New("username cannot be null or empty")
	ErrorDataNotFound     = errors.New("data not found")
	ErrorInvalidSnippetID = errors.New("invalid snippet id")
	ErrorInvalidCommentID = errors.New("invalid comment id")
	ErrorEmptyComment     = errors.New("comment cannot be empty")
	ErrorEmptySnippet     = errors.New("snippet cannot be empty")
	ErrorEmptyTitle       = errors.New("title cannot be empty")
)
