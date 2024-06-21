package services

import "graphql-comments/internal/models"

type PostService interface {
	CreatePost(title, content string, commentsEnabled bool) (error, *models.Post)
	GetPosts() (error, []*models.Post)
	GetPostByID(id string) (error, *models.Post)
}

type CommentService interface {
	CreateComment(postID string, parentID *string, content string) (error, *models.Comment)
	GetCommentsByPostID(postID string) (error, []*models.Comment)
}
