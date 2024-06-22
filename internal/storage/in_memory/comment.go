package in_memory

import (
	"graphql-comments/internal/errors"
	"graphql-comments/internal/models"
	"graphql-comments/internal/storage"
	"sort"
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

func (ct *commentRepository) GetComments(postID string, limit int, offset int) ([]*models.Comment, error) {
	ct.commentMutex.RLock()
	defer ct.commentMutex.RUnlock()
	filteredComments := make([]*models.Comment, 0, 10)
	for _, comment := range ct.comments {
		if comment.PostID == postID && comment.ParentID == nil {
			filteredComments = append(filteredComments, comment)
		}
	}

	if len(filteredComments) == 0 {
		return nil, errors.ErrCommentsNotFound
	}

	sort.Slice(filteredComments, func(i, j int) bool {
		return filteredComments[i].CreatedAt.Before(filteredComments[j].CreatedAt)
	})

	start := offset
	end := offset + limit

	if start > len(filteredComments) {
		start = len(filteredComments)
	}
	if end > len(filteredComments) {
		end = len(filteredComments)
	}

	return filteredComments[start:end], nil
}
