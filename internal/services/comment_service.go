package services

import (
	"graphql-comments/internal/errors"
	"graphql-comments/internal/models"
	"graphql-comments/internal/storage"
	"sync"
)

type commentService struct {
	commentRepo storage.CommentRepository
	postRepo    storage.PostRepository
	subscribers map[string][]chan *models.Comment
	mu          sync.RWMutex
}

func NewCommentService(comRepo storage.CommentRepository, postRep storage.PostRepository) CommentService {
	return &commentService{
		commentRepo: comRepo,
		postRepo:    postRep,
		subscribers: make(map[string][]chan *models.Comment),
	}
}

func (s *commentService) CreateComment(postID string, parentID *string, content string) (*models.Comment, error) {
	post, err := s.postRepo.GetPost(postID)
	if err != nil {
		return nil, err
	}
	if !post.CommentsEnabled {
		return nil, errors.ErrCommentsEnabled
	}
	if len(content) > 2000 {
		return nil, errors.ErrOutOfLength
	}
	comment := models.NewComment(postID, parentID, content)
	com, err := s.commentRepo.AddComment(comment)
	if err != nil {
		return nil, err
	}

	s.notifyCommentAdded(comment)

	return &com, nil
}

func (s *commentService) GetCommentsByPostID(postID string, limit int, offset int) ([]*models.Comment, error) {
	return s.commentRepo.GetComments(postID, limit, offset)
}

func (s *commentService) SubscribeToComments(postID string) (<-chan *models.Comment, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	ch := make(chan *models.Comment, 1)
	s.subscribers[postID] = append(s.subscribers[postID], ch)

	return ch, nil
}

func (s *commentService) notifyCommentAdded(comment *models.Comment) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	subscribers := s.subscribers[comment.PostID]
	for _, ch := range subscribers {
		go func(ch chan<- *models.Comment) {
			select {
			case ch <- comment:
			default:
			}
		}(ch)
	}
}
