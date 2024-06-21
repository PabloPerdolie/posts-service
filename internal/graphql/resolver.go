package graphql

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"graphql-comments/internal/models"
	"graphql-comments/internal/services"
)

type Resolver struct {
	PostService    services.PostService
	CommentService services.CommentService
}

type Config struct {
	Resolvers *Resolver
}

func NewExecutableSchema(cfg Config) graphql.ExecutableSchema {
	return NewExecutableSchema(Config{Resolvers: cfg.Resolvers})
}

func (r *Resolver) Posts(ctx context.Context) (error, []*models.Post) {
	return r.PostService.GetPosts()
}

func (r *Resolver) Post(ctx context.Context, id string) (error, *models.Post) {
	return r.PostService.GetPostByID(id)
}

func (r *Resolver) Comments(ctx context.Context, postId string) (error, []*models.Comment) {
	return r.CommentService.GetCommentsByPostID(postId)
}

func (r *Resolver) CreatePost(ctx context.Context, title string, content string, commentsEnabled bool) (error, *models.Post) {
	return r.PostService.CreatePost(title, content, commentsEnabled)
}

func (r *Resolver) CreateComment(ctx context.Context, postId string, parentId *string, content string) (error, *models.Comment) {
	return r.CommentService.CreateComment(postId, parentId, content)
}
