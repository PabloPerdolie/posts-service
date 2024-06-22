package services

import "graphql-comments/internal/models"

type PostService interface {
	CreatePost(title, content string, commentsEnabled bool) (*models.Post, error)
	GetPosts() ([]*models.Post, error)
	GetPostByID(id string) (*models.Post, error)
	ManageComments(postID string, enable bool) (*models.Post, error)
}

type CommentService interface {
	CreateComment(postID string, parentID *string, content string) (*models.Comment, error)
	GetCommentsByPostID(postID string, limit int, offset int) ([]*models.Comment, error)
	SubscribeToComments(postID string) (<-chan *models.Comment, error)
	notifyCommentAdded(comment *models.Comment)
}
