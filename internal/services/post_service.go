package services

import (
	"graphql-comments/internal/models"
	"graphql-comments/internal/storage"
)

type postService struct {
	postRepo storage.PostRepository
}

func NewPostService(repo storage.PostRepository) PostService {
	return &postService{
		postRepo: repo,
	}
}

func (s *postService) CreatePost(title, content string, commentsEnabled bool) (*models.Post, error) {
	post := models.NewPost(title, content, commentsEnabled)
	err, p := s.postRepo.AddPost(post)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (s *postService) GetPosts() (error, []*models.Post) {
	return s.postRepo.GetPosts()
}

func (s *postService) GetPostByID(id string) (error, *models.Post) {
	return s.postRepo.GetPost(id)
}
