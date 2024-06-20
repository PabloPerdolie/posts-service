package services

import "graphql-comments/internal/models"

type PostService interface {
	CreatePost(title, content string, commentsEnabled bool) (*models.Post, error)
	GetPosts() (error, []*models.Post)
	GetPostByID(id string) (error, *models.Post)
}

type CommentService interface {
}
