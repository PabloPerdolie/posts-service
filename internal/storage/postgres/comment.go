package postgres

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"graphql-comments/internal/models"
	"graphql-comments/internal/storage"
)

type CommentRepository struct {
	db *sqlx.DB
}

func NewCommentRepository(db *sqlx.DB) storage.CommentRepository {
	return &CommentRepository{
		db: db,
	}
}

func (c *CommentRepository) AddComment(comment *models.Comment) (*models.Comment, error) {
	query := `
		INSERT INTO comments (id, post_id, parent_id, content, created_at)
		VALUES ($1, $2, $3, $4, NOW())
		RETURNING id, post_id, parent_id, content, created_at`
	err := c.db.QueryRow(query, comment.ID, comment.PostID, comment.ParentID, comment.Content).
		Scan(&comment.ID, &comment.PostID, &comment.ParentID, &comment.Content, &comment.CreatedAt)
	if err != nil {
		return nil, err
	}
	return comment, nil
}

func (c *CommentRepository) GetComments(postID string, limit int, offset int) ([]*models.Comment, error) {
	query := `
		SELECT id, post_id, parent_id, content, created_at
		FROM comments
		WHERE post_id = $1 AND parent_id IS NULL
		ORDER BY created_at
		LIMIT $2 OFFSET $3`
	rows, err := c.db.Query(query, postID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []*models.Comment
	for rows.Next() {
		var comment models.Comment
		if err := rows.Scan(&comment.ID, &comment.PostID, &comment.ParentID, &comment.Content, &comment.CreatedAt); err != nil {
			return nil, err
		}
		comments = append(comments, &comment)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	for _, comment := range comments {
		replies, err := c.getReplies(comment.ID)
		if err != nil {
			return nil, err
		}
		comment.Children = replies
	}

	if len(comments) == 0 {
		return nil, errors.New("comments not found")
	}

	return comments, nil
}

func (c *CommentRepository) getReplies(parentID string) ([]*models.Comment, error) {
	query := `
		SELECT id, post_id, parent_id, content, created_at
		FROM comments
		WHERE parent_id = $1
		ORDER BY created_at`
	rows, err := c.db.Query(query, parentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var replies []*models.Comment
	for rows.Next() {
		var reply models.Comment
		if err := rows.Scan(&reply.ID, &reply.PostID, &reply.ParentID, &reply.Content, &reply.CreatedAt); err != nil {
			return nil, err
		}
		reply.Children, err = c.getReplies(reply.ID)
		if err != nil {
			return nil, err
		}
		replies = append(replies, &reply)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return replies, nil
}
