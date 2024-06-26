package storage

import "graphql-comments/internal/models"

type PostRepository interface {
	AddPost(post *models.Post) (*models.Post, error)
	GetPosts() ([]*models.Post, error)
	GetPost(id string) (*models.Post, error)
	ManageComments(postID string, enable bool) (*models.Post, error)
}

type CommentRepository interface {
	AddComment(comment *models.Comment) (*models.Comment, error)
	GetComments(postID string, limit int, offset int) ([]*models.Comment, error)
}
