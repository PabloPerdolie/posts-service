package postgres

import (
	"database/sql"
	"graphql-comments/internal/errors"

	"github.com/jmoiron/sqlx"
	"graphql-comments/internal/models"
	"graphql-comments/internal/storage"
)

type PostRepository struct {
	db *sqlx.DB
}

func NewPostRepository(db *sqlx.DB) storage.PostRepository {
	return &PostRepository{
		db: db,
	}
}

func (p *PostRepository) AddPost(post *models.Post) (*models.Post, error) {
	query := `
		INSERT INTO posts (id, title, content, comments_enabled, created_at)
		VALUES ($1, $2, $3, $4, NOW())
		RETURNING id, title, content, comments_enabled, created_at`
	err := p.db.QueryRow(query, post.ID, post.Title, post.Content, post.CommentsEnabled).
		Scan(&post.ID, &post.Title, &post.Content, &post.CommentsEnabled, &post.CreatedAt)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (p *PostRepository) GetPosts() ([]*models.Post, error) {
	query := `
		SELECT id, title, content, comments_enabled, created_at
		FROM posts`
	rows, err := p.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []*models.Post
	for rows.Next() {
		var post models.Post
		if err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.CommentsEnabled, &post.CreatedAt); err != nil {
			return nil, err
		}
		posts = append(posts, &post)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if len(posts) == 0 {
		return nil, errors.ErrPostNotFound
	}

	return posts, nil
}

func (p *PostRepository) GetPost(id string) (*models.Post, error) {
	query := `
		SELECT id, title, content, comments_enabled, created_at
		FROM posts
		WHERE id = $1`
	var post models.Post
	err := p.db.QueryRow(query, id).Scan(&post.ID, &post.Title, &post.Content, &post.CommentsEnabled, &post.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.ErrPostNotFound
		}
		return nil, err
	}
	return &post, nil
}

func (p *PostRepository) ManageComments(postID string, enable bool) (*models.Post, error) {
	query := `
		UPDATE posts
		SET comments_enabled = $1
		WHERE id = $2
		RETURNING id, title, content, comments_enabled, created_at`
	var post models.Post
	err := p.db.QueryRow(query, enable, postID).
		Scan(&post.ID, &post.Title, &post.Content, &post.CommentsEnabled, &post.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.ErrPostNotFound
		}
		return nil, err
	}
	return &post, nil
}
