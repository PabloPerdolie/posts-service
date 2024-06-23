package postgres

import (
	"github.com/jmoiron/sqlx"
	"graphql-comments/internal/models"
	"graphql-comments/internal/storage"
)

type PostRepository struct {
	db *sqlx.DB
}

func NewPostRepository(db *sqlx.DB) storage.PostRepository {
	return &PostRepository{
		db: db,
	}
}

func (p PostRepository) AddPost(post *models.Post) (models.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (p PostRepository) GetPosts() ([]*models.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (p PostRepository) GetPost(id string) (*models.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (p PostRepository) ManageComments(postID string, enable bool) (*models.Post, error) {
	//TODO implement me
	panic("implement me")
}
