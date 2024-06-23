package app

import (
	"graphql-comments/internal/config"
	"graphql-comments/internal/services"
	"graphql-comments/internal/storage"
	"graphql-comments/internal/storage/in_memory"
	"graphql-comments/internal/storage/postgres"
	"graphql-comments/pkg/client"
)

type serviceProvider struct {
	commentRepo    storage.CommentRepository
	commentService services.CommentService
	postRepo       storage.PostRepository
	postService    services.PostService
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (sp *serviceProvider) postgresRepo() error {
	dbConfig := config.CONFIG.DB
	db, err := client.InitDB(dbConfig.User, dbConfig.Password, dbConfig.Name, dbConfig.Host, dbConfig.Port)
	if err != nil {
		return err
	}
	sp.postRepo = postgres.NewPostRepository(db)
	sp.commentRepo = postgres.NewCommentRepository(db, sp.postRepo)
	return nil
}

func (sp *serviceProvider) inMemoryRepo() {
	sp.commentRepo = in_memory.NewCommentRepository()
	sp.postRepo = in_memory.NewPostRepository()
}

func (sp *serviceProvider) initServices(useInMemory bool) error {
	if useInMemory {
		sp.inMemoryRepo()
	} else {
		err := sp.postgresRepo()
		if err != nil {
			return err
		}
	}
	sp.commentService = services.NewCommentService(sp.commentRepo, sp.postRepo)
	sp.postService = services.NewPostService(sp.postRepo)
	return nil
}
