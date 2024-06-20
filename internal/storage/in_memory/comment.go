package in_memory

import "graphql-comments/internal/models"

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
