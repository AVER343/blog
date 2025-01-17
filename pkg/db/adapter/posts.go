package adapter

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/aver343/blog/pkg/db/sqlc"
	"github.com/aver343/blog/pkg/models"
)

var (
	ErrNotFound = errors.New("The requested data is not present")
)

type SQLCPostRepository struct {
	query *sqlc.Queries
}

func NewSQLCPostRepository(queries *sqlc.Queries) *SQLCPostRepository {
	return &SQLCPostRepository{query: queries}
}

func (p *SQLCPostRepository) Create(ctx context.Context, m *models.Post) error {
	postParam := &sqlc.CreatePostParams{
		Title:   m.Title,
		Content: m.Content,
		UserID:  m.UserID,
	}

	savedPost, err := p.query.CreatePost(ctx, postParam)
	if err != nil {
		return err
	}
	m.ID = savedPost.ID
	m.CreatedAt = savedPost.CreatedAt.Time.String()
	m.UpdatedAt = savedPost.UpdatedAt.Time.String()
	return nil
}

func (p *SQLCPostRepository) GetPostByID(ctx context.Context, id int64) (*models.Post, error) {

	post, err := p.query.GetPostByID(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrNotFound
		}
		return nil, err
	}
	fmt.Print(post)
	modelPost := models.NewPost(post.ID, post.Content, post.Title, post.UserID, nil)
	return &modelPost, nil
}

func (p *SQLCPostRepository) PatchByID(ctx context.Context, id int64) (*models.Post, error) {

	post, err := p.query.GetPostByID(ctx, id)
	if err != nil {
		return nil, err
	}
	fmt.Print(post)
	modelPost := models.NewPost(post.ID, post.Content, post.Title, post.UserID, nil)
	return &modelPost, nil
}
