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
		WHERE post_id = $1
		ORDER BY created_at
		LIMIT $2 OFFSET $3`
	rows, err := c.db.Query(query, postID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	//тут постоянное увеличение слайса и мапы, хз как исправить
	//тк при обычном переборе через rows.Next() указатель на текущую останется в конце и этот цикл тупо скипнется,
	//СУПЕР не оптимизированное решение
	//так нельзя делать, извините
	commentsMap := make(map[string]*models.Comment)
	var comments []*models.Comment
	for rows.Next() {
		var comment models.Comment
		if err := rows.Scan(&comment.ID, &comment.PostID, &comment.ParentID, &comment.Content, &comment.CreatedAt); err != nil {
			return nil, err
		}
		commentsMap[comment.ID] = &comment
		if comment.ParentID == nil {
			comments = append(comments, &comment)
		} else {
			parentComment, exists := commentsMap[*comment.ParentID]
			if exists {
				parentComment.Children = append(parentComment.Children, &comment)
			}
		}

	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if len(comments) == 0 {
		return nil, errors.New("comments not found")
	}

	return comments, nil
}
