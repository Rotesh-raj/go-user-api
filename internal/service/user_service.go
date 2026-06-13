package service

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/user/go-user-api/db/sqlc"
	"github.com/user/go-user-api/internal/models"
	"github.com/user/go-user-api/internal/repository"
	"github.com/user/go-user-api/pkg/utils"
)

type UserService interface {
	CreateUser(ctx context.Context, req *models.CreateUserRequest) (*models.User, error)
	GetUser(ctx context.Context, id int32) (*models.User, error)
	ListUsers(ctx context.Context, page, limit int32) ([]*models.User, error)
	UpdateUser(ctx context.Context, id int32, req *models.UpdateUserRequest) (*models.User, error)
	DeleteUser(ctx context.Context, id int32) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(ctx context.Context, req *models.CreateUserRequest) (*models.User, error) {
	dob, _ := time.Parse("2006-01-02", req.Dob)
	pgDob := pgtype.Date{Time: dob, Valid: true}

	user, err := s.repo.Create(ctx, sqlc.CreateUserParams{
		Name: req.Name,
		Dob:  pgDob,
	})
	if err != nil {
		return nil, err
	}

	return &models.User{
		ID:   user.ID,
		Name: user.Name,
		Dob:  user.Dob.Time,
	}, nil
}

func (s *userService) GetUser(ctx context.Context, id int32) (*models.User, error) {
	user, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &models.User{
		ID:   user.ID,
		Name: user.Name,
		Dob:  user.Dob.Time,
		Age:  utils.CalculateAge(user.Dob.Time),
	}, nil
}

func (s *userService) ListUsers(ctx context.Context, page, limit int32) ([]*models.User, error) {
	offset := (page - 1) * limit
	users, err := s.repo.List(ctx, sqlc.ListUsersParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, err
	}

	var result []*models.User
	for _, u := range users {
		result = append(result, &models.User{
			ID:   u.ID,
			Name: u.Name,
			Dob:  u.Dob.Time,
			Age:  utils.CalculateAge(u.Dob.Time),
		})
	}
	return result, nil
}

func (s *userService) UpdateUser(ctx context.Context, id int32, req *models.UpdateUserRequest) (*models.User, error) {
	dob, _ := time.Parse("2006-01-02", req.Dob)
	pgDob := pgtype.Date{Time: dob, Valid: true}

	user, err := s.repo.Update(ctx, sqlc.UpdateUserParams{
		ID:   id,
		Name: req.Name,
		Dob:  pgDob,
	})
	if err != nil {
		return nil, err
	}

	return &models.User{
		ID:   user.ID,
		Name: user.Name,
		Dob:  user.Dob.Time,
	}, nil
}

func (s *userService) DeleteUser(ctx context.Context, id int32) error {
	return s.repo.Delete(ctx, id)
}
