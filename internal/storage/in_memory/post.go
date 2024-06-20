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

func (st *postRepository) AddPost(post *models.Post) (error, models.Post) {
	st.postMutex.Lock()
	defer st.postMutex.Unlock()
	st.posts[post.ID] = post
	return nil, *post
}

func (st *postRepository) GetPosts() (error, []*models.Post) {
	st.postMutex.RLock()
	defer st.postMutex.RUnlock()
	posts := make([]*models.Post, 0, len(st.posts))
	for _, post := range st.posts {
		posts = append(posts, post)
	}
	if len(posts) == 0 {
		return errors.ErrPostNotFound, nil
	}
	return nil, posts
}

func (st *postRepository) GetPost(id string) (error, *models.Post) {
	st.postMutex.RLock()
	defer st.postMutex.RUnlock()
	post, exists := st.posts[id]
	if !exists {
		return errors.ErrPostNotFound, nil
	}
	return nil, post
}
