package adapter

import (
	"context"

	"github.com/aver343/blog/pkg/db/sqlc"
	"github.com/aver343/blog/pkg/models"
)

type SQLCUserRepository struct {
	queries *sqlc.Queries
}

func NewSQLCUserRepository(queries *sqlc.Queries) *SQLCUserRepository {
	return &SQLCUserRepository{queries: queries}
}

func (r *SQLCUserRepository) Create(ctx context.Context, user *models.User) error {
	params := &sqlc.CreateUserParams{
		Email:    user.Email,
		Username: user.Username,
		Password: user.Password,
	}
	createdUser, err := r.queries.CreateUser(ctx, params)
	if err != nil {
		return err
	}
	user.ID = createdUser.ID
	user.CreatedAt = createdUser.CreatedAt.Time.String()
	user.UpdatedAt = createdUser.UpdatedAt.Time.String()

	return nil
}
