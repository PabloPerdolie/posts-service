package in_memory

import (
	"graphql-comments/internal/models"
	"graphql-comments/internal/storage"
	"sync"
)

type commentRepository struct {
	comments     map[string]*models.Comment
	commentMutex sync.RWMutex
}

func NewCommentRepository() storage.CommentRepository {
	return &commentRepository{
		comments: make(map[string]*models.Comment, 10),
	}
}

func (ct *commentRepository) AddComment(comment *models.Comment) (models.Comment, error) {
	ct.commentMutex.Lock()
	defer ct.commentMutex.Unlock()
	ct.comments[comment.ID] = comment
	if comment.ParentID != nil {
		parent, exists := ct.comments[*comment.ParentID]
		if exists {
			parent.Children = append(parent.Children, comment)
		}
	}
	return *comment, nil
}

func (ct *commentRepository) GetComments(postID string) ([]*models.Comment, error) {
	ct.commentMutex.RLock()
	defer ct.commentMutex.RUnlock()
	comments := make([]*models.Comment, 255)
	for _, comment := range ct.comments {
		if comment.PostID == postID && comment.ParentID == nil {
			comments = append(comments, comment)
		}
	}
	return comments, nil
}
