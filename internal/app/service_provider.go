package app

import (
	"graphql-comments/internal/services"
	"graphql-comments/internal/storage"
	"graphql-comments/internal/storage/in_memory"
)

type serviceProvider struct {
	commentRepo    storage.CommentRepository
	commentService services.CommentService
	postRepo       storage.PostRepository
	postService    services.PostService
}

func newServiceProvider(useInMemory bool) *serviceProvider {
	sp := &serviceProvider{}

	if useInMemory {
		sp.commentRepo = in_memory.NewCommentRepository()
		sp.postRepo = in_memory.NewPostRepository()
	} else {
		//sp.commentRepo = db.NewCommentRepository()
		//sp.postRepo = db.NewPostRepository()
	}

	sp.commentService = services.NewCommentService(sp.commentRepo)
	sp.postService = services.NewPostService(sp.postRepo)

	return sp
}
