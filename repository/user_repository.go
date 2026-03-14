package repository

import (
	"context"
	"errors"

	"github.com/sidz111/jwt-twitter-msg/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUser(ctx context.Context, id uint) (*models.User, error)
	GetAllUsers(ctx context.Context) ([]*models.User, error)
	UpdateUser(ctx context.Context, user *models.User) error
	DeleteUser(ctx context.Context, id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(ctx context.Context, user *models.User) error {
	result := r.db.WithContext(ctx).Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (r *userRepository) GetUser(ctx context.Context, id uint) (*models.User, error) {
	var user models.User
	result := r.db.Preload("Posts").WithContext(ctx).First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
func (r *userRepository) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	var users []*models.User
	result := r.db.Preload("Posts").WithContext(ctx).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}
func (r *userRepository) UpdateUser(ctx context.Context, user *models.User) error {
	// result := r.db.WithContext(ctx).Save(user)
	// if result.Error != nil {
	// 	return result.Error
	// }
	result := r.db.WithContext(ctx).Model(&models.User{}).Where("id = ?", user.ID).Updates(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (r *userRepository) DeleteUser(ctx context.Context, id uint) error {
	result := r.db.WithContext(ctx).Delete(&models.User{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected <= 0 {
		return errors.New("user not found")
	}
	return nil
}
