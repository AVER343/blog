package adapter

import (
	"context"
	"database/sql"
	"errors"

	"github.com/aver343/blog/pkg/db/dto"
	"github.com/aver343/blog/pkg/db/sqlc"
	"github.com/aver343/blog/pkg/models"
)

var (
	ErrNotFound = errors.New("the requested data is not present")
)

type SQLCPostRepository struct {
	query *sqlc.Queries
}

func NewSQLCPostRepository(queries *sqlc.Queries) *SQLCPostRepository {
	return &SQLCPostRepository{query: queries}
}

func (p *SQLCPostRepository) Create(ctx context.Context, m *dto.CreatePostPayload) (*models.Post, error) {
	postParam := &sqlc.CreatePostParams{
		Title:   m.Title,
		Content: m.Content,
		UserID:  m.UserID,
	}

	dbPost, err := p.query.CreatePost(ctx, postParam)
	if err != nil {
		return nil, err
	}
	post := models.NewPost(dbPost.ID, dbPost.Content, dbPost.Title, dbPost.UserID, dbPost.Tags)
	return post, nil
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
	modelPost := models.NewPost(post.ID, post.Content, post.Title, post.UserID, nil)
	return modelPost, nil
}

func (p *SQLCPostRepository) PatchByID(ctx context.Context, id int64) (*models.Post, error) {

	post, err := p.query.GetPostByID(ctx, id)
	if err != nil {
		return nil, err
	}
	modelPost := models.NewPost(post.ID, post.Content, post.Title, post.UserID, nil)
	return modelPost, nil
}

func (p *SQLCPostRepository) GetAllPosts(ctx context.Context) ([]*models.Post, error) {
	var posts []*models.Post = make([]*models.Post, 0)
	dbPosts, err := p.query.GetAllPosts(ctx, nil)
	if err != nil {
		return nil, err
	}
	for _, elem := range dbPosts {
		post := models.NewPost(elem.ID, elem.Content, elem.Title, elem.UserID, elem.Tags)
		posts = append(posts, post)
	}
	return posts, nil
}

func (p *SQLCPostRepository) GetPostByUserID(ctx context.Context, payload *dto.GetPostByUserIDPayload) ([]*models.Post, error) {
	var posts []*models.Post = make([]*models.Post, 0)
	dbPosts, err := p.query.GetPostByUserID(ctx, payload.UserID)
	if err != nil {
		return nil, err
	}
	for _, elem := range dbPosts {
		post := models.NewPost(elem.ID, elem.Content, elem.Title, elem.UserID, elem.Tags)
		posts = append(posts, post)
	}
	return posts, nil
}
