package adapter

import (
	"context"

	"github.com/aver343/blog/pkg/db/sqlc"
	"github.com/aver343/blog/pkg/models"
)

type SQLCUserRepository struct {
	query *sqlc.Queries
}

func NewSQLCUserRepository(queries *sqlc.Queries) *SQLCUserRepository {
	return &SQLCUserRepository{query: queries}
}

func (r *SQLCUserRepository) Create(ctx context.Context, user *models.User) error {
	params := &sqlc.CreateUserParams{
		Email:    user.Email,
		Username: user.Username,
		Password: user.Password,
	}
	createdUser, err := r.query.CreateUser(ctx, params)
	if err != nil {
		return err
	}
	user.ID = createdUser.ID
	user.CreatedAt = createdUser.CreatedAt.Time.String()
	user.UpdatedAt = createdUser.UpdatedAt.Time.String()

	return nil
}

func (r *SQLCUserRepository) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	var users []*models.User = make([]*models.User, 0)
	dbUserData, err := r.query.GetAllUsers(ctx, 2)
	if err != nil {
		return nil, err
	}
	for _, elem := range dbUserData {
		user := models.NewUser(elem.ID, elem.Username, elem.Password, elem.Email)
		users = append(users, user)
	}
	return users, nil
}

func (r *SQLCUserRepository) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	user, err := r.query.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	modelUser := models.NewUser(user.ID, user.Username, user.Password, user.Email)
	return modelUser, nil
}
