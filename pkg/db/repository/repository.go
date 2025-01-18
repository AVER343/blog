package repository

import (
	"context"
	"database/sql"

	"github.com/aver343/blog/pkg/db/adapter"
	"github.com/aver343/blog/pkg/db/sqlc"
	"github.com/aver343/blog/pkg/models"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	GetUserByID(ctx context.Context, ID string) (*models.User, error)
	GetAllUsers(ctx context.Context) ([]*models.User, error)
}
type PostRepository interface {
	Create(ctx context.Context, post *models.Post) error
	GetPostByID(ctx context.Context, ID int64) (*models.Post, error)
	GetAllPosts(ctx context.Context) ([]*models.Post, error)
}

type Repository struct {
	Post PostRepository
	User UserRepository
}

func NewRepository(db *sql.DB) Repository {
	sqlc := sqlc.New(db)
	userRepo := adapter.NewSQLCUserRepository(sqlc)
	postRepo := adapter.NewSQLCPostRepository(sqlc)
	return Repository{
		User: userRepo,
		Post: postRepo,
	}
}
