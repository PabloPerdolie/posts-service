package in_memory

import (
	"graphql-comments/internal/models"
	"graphql-comments/internal/storage"
	"sync"
)

type InMemoryStore struct {
	posts        map[string]*models.Post
	comments     map[string]*models.Comment
	postMutex    sync.RWMutex
	commentMutex sync.RWMutex
}

func NewInMemoryStore() storage.Storage {
	return &InMemoryStore{
		posts:    make(map[string]*models.Post, 10),
		comments: make(map[string]*models.Comment, 20),
	}
}

func (st *InMemoryStore) AddPost(post *models.Post) {
	st.postMutex.Lock()
	defer st.postMutex.Unlock()
	st.posts[post.ID] = post
}

func (st *InMemoryStore) GetPosts() []*models.Post {
	st.postMutex.RLock()
	defer st.postMutex.RUnlock()
	posts := make([]*models.Post, 0, len(st.posts))
	for _, post := range st.posts {
		posts = append(posts, post)
	}
	return posts
}

func (st *InMemoryStore) GetPost(id string) (*models.Post, bool) {
	st.postMutex.RLock()
	defer st.postMutex.RUnlock()
	post, exists := st.posts[id]
	return post, exists
}

func (st *InMemoryStore) AddComment(comment *models.Comment) {
	st.commentMutex.Lock()
	defer st.commentMutex.Unlock()
	st.comments[comment.ID] = comment
	if comment.ParentID != nil {
		parent, exists := st.comments[*comment.ParentID]
		if exists {
			parent.Children = append(parent.Children, comment)
		}
	} else {
		post, exists := st.posts[comment.PostID]
		if exists {
			post.Comments = append(post.Comments, comment)
		}
	}
}

func (st *InMemoryStore) GetComments(postID string) []*models.Comment {
	st.commentMutex.RLock()
	defer st.commentMutex.RUnlock()
	comments := make([]*models.Comment, 0)
	for _, comment := range st.comments {
		if comment.PostID == postID && comment.ParentID == nil {
			comments = append(comments, comment)
		}
	}
	return comments
}
