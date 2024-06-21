package services

import (
	"graphql-comments/internal/errors"
	"graphql-comments/internal/models"
	"graphql-comments/internal/storage"
)

type commentService struct {
	commentRepo storage.CommentRepository
}

func NewCommentService(repo storage.CommentRepository) CommentService {
	return &commentService{
		commentRepo: repo,
	}
}

func (s *commentService) CreateComment(postID string, parentID *string, content string) (*models.Comment, error) {
	if len(content) > 2000 {
		return nil, errors.ErrOutOfLength
	}
	comment := models.NewComment(postID, parentID, content)
	com, err := s.commentRepo.AddComment(comment)
	if err != nil {
		return nil, err
	}
	return &com, nil
}

func (s *commentService) GetCommentsByPostID(postID string) ([]*models.Comment, error) {
	return s.commentRepo.GetComments(postID)
}
