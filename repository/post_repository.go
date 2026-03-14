package repository

import (
	"context"
	"errors"

	"github.com/sidz111/jwt-twitter-msg/models"
	"gorm.io/gorm"
)

type PostRepository interface {
	CreatePost(ctx context.Context, post *models.Post) error
	GetPost(ctx context.Context, id uint) (*models.Post, error)
	GetAllPosts(ctx context.Context) ([]*models.Post, error)
	UpdatePost(ctx context.Context, post *models.Post) error
	DeletePost(ctx context.Context, id uint) error
	GetPostsByUserId(ctx context.Context, userID uint) ([]*models.Post, error)
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{db: db}
}

func (r *postRepository) CreatePost(ctx context.Context, post *models.Post) error {
	return r.db.WithContext(ctx).Create(post).Error
}
func (r *postRepository) GetPost(ctx context.Context, id uint) (*models.Post, error) {
	var post models.Post
	result := r.db.WithContext(ctx).Preload("User").First(&post, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &post, nil
}
func (r *postRepository) GetAllPosts(ctx context.Context) ([]*models.Post, error) {
	var posts []*models.Post
	result := r.db.WithContext(ctx).Preload("User").Find(&posts)
	if result.Error != nil {
		return nil, result.Error
	}
	return posts, nil
}
func (r *postRepository) UpdatePost(ctx context.Context, post *models.Post) error {
	result := r.db.WithContext(ctx).Model(&models.Post{}).Where("id =?", post.ID).Updates(post)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (r *postRepository) DeletePost(ctx context.Context, id uint) error {
	result := r.db.WithContext(ctx).Delete(&models.Post{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("post not found")
	}
	return nil
}

func (r *postRepository) GetPostsByUserId(ctx context.Context, userID uint) ([]*models.Post, error) {
	var posts []*models.Post
	result := r.db.WithContext(ctx).Model(&models.Post{}).Where("id =?", userID).Find(&posts)
	if result.Error != nil {
		return nil, result.Error
	}
	return posts, nil
}
