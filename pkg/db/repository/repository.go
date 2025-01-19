package repository

import (
	"context"
	"database/sql"

	"github.com/aver343/blog/pkg/db/adapter"
	"github.com/aver343/blog/pkg/db/dto"
	"github.com/aver343/blog/pkg/db/sqlc"
	"github.com/aver343/blog/pkg/models"
)

type UserRepository interface {
	Create(context.Context, *dto.RegisterUserPayload) (*models.User, error)
	GetUserByID(context.Context, string) (*models.User, error)
	GetAllUsers(context.Context) ([]*models.User, error)
}
type PostRepository interface {
	Create(context.Context, *dto.CreatePostPayload) (*models.Post, error)
	GetPostByID(context.Context, int64) (*models.Post, error)
	GetAllPosts(context.Context) ([]*models.Post, error)
	GetPostByUserID(context.Context, *dto.GetPostByUserIDPayload) ([]*models.Post, error)
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
