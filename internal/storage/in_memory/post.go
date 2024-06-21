package in_memory

import (
	"graphql-comments/internal/errors"
	"graphql-comments/internal/models"
	"graphql-comments/internal/storage"
	"sync"
)

type postRepository struct {
	posts     map[string]*models.Post
	postMutex sync.RWMutex
}

func NewPostRepository() storage.PostRepository {
	return &postRepository{
		posts: make(map[string]*models.Post, 10),
	}
}

func (st *postRepository) AddPost(post *models.Post) (models.Post, error) {
	st.postMutex.Lock()
	defer st.postMutex.Unlock()
	st.posts[post.ID] = post
	return *post, nil
}

func (st *postRepository) GetPosts() ([]*models.Post, error) {
	st.postMutex.RLock()
	defer st.postMutex.RUnlock()
	posts := make([]*models.Post, 0, len(st.posts))
	for _, post := range st.posts {
		posts = append(posts, post)
	}
	if len(posts) == 0 {
		return nil, errors.ErrPostNotFound
	}
	return posts, nil
}

func (st *postRepository) GetPost(id string) (*models.Post, error) {
	st.postMutex.RLock()
	defer st.postMutex.RUnlock()
	post, exists := st.posts[id]
	if !exists {
		return nil, errors.ErrPostNotFound
	}
	return post, nil
}
