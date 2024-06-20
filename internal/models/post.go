package models

import "github.com/google/uuid"

type Post struct {
	ID              string
	Title           string
	Content         string
	Comments        []*Comment
	CommentsEnabled bool
}

func NewPost(title, content string, commentsEnabled bool) *Post {
	return &Post{
		ID:              uuid.New().String(),
		Title:           title,
		Content:         content,
		CommentsEnabled: commentsEnabled,
	}
}
