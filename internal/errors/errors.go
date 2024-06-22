package errors

import "errors"

var (
	ErrPostNotFound     = errors.New("post not found")
	ErrCommentsNotFound = errors.New("comments not found")
	ErrOutOfLength      = errors.New("content length exceeds 2000 characters")
	ErrCommentsEnabled  = errors.New("comments for this posts enabled")
)
