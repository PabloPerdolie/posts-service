package postgres

import (
	"github.com/jmoiron/sqlx"
	"graphql-comments/internal/models"
	"graphql-comments/internal/storage"
)

type CommentRepository struct {
	postRepo storage.PostRepository
	db       *sqlx.DB
}

func NewCommentRepository(db *sqlx.DB, postRepo storage.PostRepository) storage.CommentRepository {
	return &CommentRepository{
		db:       db,
		postRepo: postRepo,
	}
}

func (c CommentRepository) AddComment(comment *models.Comment) (models.Comment, error) {
	//TODO implement me
	panic("implement me")
}

func (c CommentRepository) GetComments(postID string, limit int, offset int) ([]*models.Comment, error) {
	//TODO implement me
	panic("implement me")
}
