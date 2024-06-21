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

func (s *commentService) CreateComment(postID string, parentID *string, content string) (error, *models.Comment) {
	if len(content) > 2000 {
		return errors.ErrOutOfLength, nil
	}
	comment := models.NewComment(postID, parentID, content)
	err, com := s.commentRepo.AddComment(comment)
	if err != nil {
		return err, nil
	}
	return nil, &com
}

func (s *commentService) GetCommentsByPostID(postID string) (error, []*models.Comment) {
	return s.commentRepo.GetComments(postID)
}
