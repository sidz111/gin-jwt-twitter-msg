package service

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/sidz111/jwt-twitter-msg/models"
	"github.com/sidz111/jwt-twitter-msg/repository"
)

type UserService interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUser(ctx context.Context, id uint) (*models.User, error)
	GetAllUsers(ctx context.Context) ([]*models.User, error)
	UpdateUser(ctx context.Context, user *models.User) error
	DeleteUser(ctx context.Context, id uint) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(ctx context.Context, user *models.User) error {
	if err := ValidateUser(user); err != nil {
		return err
	}
	user.UUID = uuid.NewString()
	return s.repo.CreateUser(ctx, user)
}
func (s *userService) GetUser(ctx context.Context, id uint) (*models.User, error) {
	if id <= 0 {
		return nil, errors.New("id should be positive")
	}
	return s.repo.GetUser(ctx, id)
}
func (s *userService) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	return s.repo.GetAllUsers(ctx)
}
func (s *userService) UpdateUser(ctx context.Context, user *models.User) error {
	// if err := ValidateUser(user); err != nil { // this is not mandetory to validate all fields in update
	// 	return err
	// }
	return s.repo.UpdateUser(ctx, user)
}
func (s *userService) DeleteUser(ctx context.Context, id uint) error {
	if id <= 0 {
		return errors.New("id should be positive")
	}
	return s.repo.DeleteUser(ctx, id)
}

func ValidateUser(user *models.User) error {
	if user.Bio == "" {
		return errors.New("Bio required")
	}
	if user.Email == "" {
		return errors.New("email required")
	}
	if user.Password == "" {
		return errors.New("password required")
	}
	if user.Username == "" {
		return errors.New("username required")
	}
	return nil
}
