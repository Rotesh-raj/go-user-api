package repository

import (
	"context"

	"github.com/user/go-user-api/db/sqlc"
)

type UserRepository interface {
	Create(ctx context.Context, arg sqlc.CreateUserParams) (sqlc.User, error)
	GetByID(ctx context.Context, id int32) (sqlc.User, error)
	List(ctx context.Context, arg sqlc.ListUsersParams) ([]sqlc.User, error)
	Update(ctx context.Context, arg sqlc.UpdateUserParams) (sqlc.User, error)
	Delete(ctx context.Context, id int32) error
}

type userRepository struct {
	queries *sqlc.Queries
}

func NewUserRepository(queries *sqlc.Queries) UserRepository {
	return &userRepository{queries: queries}
}

func (r *userRepository) Create(ctx context.Context, arg sqlc.CreateUserParams) (sqlc.User, error) {
	return r.queries.CreateUser(ctx, arg)
}

func (r *userRepository) GetByID(ctx context.Context, id int32) (sqlc.User, error) {
	return r.queries.GetUser(ctx, id)
}

func (r *userRepository) List(ctx context.Context, arg sqlc.ListUsersParams) ([]sqlc.User, error) {
	return r.queries.ListUsers(ctx, arg)
}

func (r *userRepository) Update(ctx context.Context, arg sqlc.UpdateUserParams) (sqlc.User, error) {
	return r.queries.UpdateUser(ctx, arg)
}

func (r *userRepository) Delete(ctx context.Context, id int32) error {
	return r.queries.DeleteUser(ctx, id)
}
