package models

import (
	"github.com/google/uuid"
	"time"
)

type Post struct {
	ID              string
	Title           string
	Content         string
	CommentsEnabled bool
	CreatedAt       time.Time
}

func NewPost(title, content string, commentsEnabled bool) *Post {
	return &Post{
		ID:              uuid.New().String(),
		Title:           title,
		Content:         content,
		CommentsEnabled: commentsEnabled,
		CreatedAt:       time.Now(),
	}
}
