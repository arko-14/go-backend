package service

import (
	"context"
	db "go-backend-task/db/sqlc"
	"go-backend-task/internal/models"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type UserService struct {
	q *db.Queries
}

func NewUserService(q *db.Queries) *UserService {
	return &UserService{q: q}
}

// Create User
func (s *UserService) CreateUser(ctx context.Context, req models.UserRequest) (models.UserResponse, error) {
	parsedDOB, _ := time.Parse("2006-01-02", req.DOB)

	user, err := s.q.CreateUser(ctx, db.CreateUserParams{
		Name: req.Name,
		Dob:  pgtype.Date{Time: parsedDOB, Valid: true},
	})

	if err != nil {
		return models.UserResponse{}, err
	}

	return models.UserResponse{
		ID:   user.ID,
		Name: user.Name,
		DOB:  user.Dob.Time.Format("2006-01-02"),
	}, nil
}

// Get User (Calculates Age)
func (s *UserService) GetUser(ctx context.Context, id int64) (models.UserResponse, error) {
	user, err := s.q.GetUser(ctx, id)
	if err != nil {
		return models.UserResponse{}, err
	}

	return models.UserResponse{
		ID:   user.ID,
		Name: user.Name,
		DOB:  user.Dob.Time.Format("2006-01-02"),
		Age:  models.CalculateAge(user.Dob.Time),
	}, nil
}

// List Users (Calculates Age for all)
func (s *UserService) ListUsers(ctx context.Context) ([]models.UserResponse, error) {
	users, err := s.q.ListUsers(ctx)
	if err != nil {
		return nil, err
	}

	var response []models.UserResponse
	for _, user := range users {
		response = append(response, models.UserResponse{
			ID:   user.ID,
			Name: user.Name,
			DOB:  user.Dob.Time.Format("2006-01-02"),
			Age:  models.CalculateAge(user.Dob.Time),
		})
	}
	return response, nil
}

// Update User
func (s *UserService) UpdateUser(ctx context.Context, id int64, req models.UserRequest) (models.UserResponse, error) {
	parsedDOB, _ := time.Parse("2006-01-02", req.DOB)

	user, err := s.q.UpdateUser(ctx, db.UpdateUserParams{
		ID:   id,
		Name: req.Name,
		Dob:  pgtype.Date{Time: parsedDOB, Valid: true},
	})

	if err != nil {
		return models.UserResponse{}, err
	}

	return models.UserResponse{
		ID:   user.ID,
		Name: user.Name,
		DOB:  user.Dob.Time.Format("2006-01-02"),
	}, nil
}

// Delete User
func (s *UserService) DeleteUser(ctx context.Context, id int64) error {
	return s.q.DeleteUser(ctx, id)
}