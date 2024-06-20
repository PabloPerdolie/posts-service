package models

import (
	"github.com/google/uuid"
	"time"
)

type Comment struct {
	ID        string
	PostID    string
	ParentID  *string
	Content   string
	CreatedAt time.Time
	Children  []*Comment
}

func NewComment(postID string, parentID *string, content string) *Comment {
	return &Comment{
		ID:        uuid.New().String(),
		PostID:    postID,
		ParentID:  parentID,
		Content:   content,
		CreatedAt: time.Now(),
	}
}
