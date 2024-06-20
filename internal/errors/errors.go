package errors

import "errors"

var (
	ErrPostNotFound     = errors.New("post not found")
	ErrCommentsNotFound = errors.New("comments not found")
)
