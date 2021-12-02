package business

import "errors"

var (
	ErrorEmptyPassword    = errors.New("password cannot be null or empty")
	ErrorInvalidEmail     = errors.New("this is not valid email, please using valid email provider and not using disposal email or spam email")
	ErrorInvalidPassword  = errors.New("wrong password")
	ErrorEmptyEmail       = errors.New("email cannot be null or empty")
	ErrorInvalidUsername  = errors.New("username cannot be null or empty")
	ErrorDataNotFound     = errors.New("data not found")
	ErrorInvalidSnippetID = errors.New("invalid snippet id")
	ErrorInvalidCommentID = errors.New("invalid comment id")
	ErrorEmptyComment     = errors.New("comment cannot be empty")
	ErrorEmptySnippet     = errors.New("snippet cannot be empty")
	ErrorEmptyTitle       = errors.New("title cannot be empty")
)
