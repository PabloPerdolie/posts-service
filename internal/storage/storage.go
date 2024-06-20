package storage

import "graphql-comments/internal/models"

type PostRepository interface {
	AddPost(post *models.Post) (error, models.Post)
	GetPosts() (error, []*models.Post)
	GetPost(id string) (error, *models.Post)
}

type CommentRepository interface {
	AddComment(comment *models.Comment) (error, models.Comment)
	GetComments(postID string) (error, []*models.Comment)
}
