package adapter

import (
	"context"

	"github.com/aver343/blog/pkg/db/sqlc"
	"github.com/aver343/blog/pkg/models"
	"github.com/google/uuid"
)

type SQLCPostRepository struct {
	query *sqlc.Queries
}

func NewSQLCPostRepository(queries *sqlc.Queries) *SQLCPostRepository {
	return &SQLCPostRepository{query: queries}
}

func (p *SQLCPostRepository) Create(ctx context.Context, m *models.Post) (error) {
	userId, _ := uuid.FromBytes([]byte(m.UserID))
	// if err != nil {
	// 	return err
	// }
	postParam := &sqlc.CreatePostParams{
		Title:   m.Title,
		Content: m.Content,
		UserID:  userId,
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

func (p *SQLCPostRepository) GetByID(ctx context.Context, id int64) (*models.Post, error) {

	post, err := p.query.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	modelPost := models.NewPost(post.Content, post.Title, post.UserID.String(), nil)
	return &modelPost, nil
}
