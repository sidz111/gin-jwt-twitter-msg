package service

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/sidz111/jwt-twitter-msg/models"
	"github.com/sidz111/jwt-twitter-msg/repository"
)

type PostService interface {
	CreatePost(ctx context.Context, post *models.Post) error
	GetPost(ctx context.Context, id uint) (*models.Post, error)
	GetAllPosts(ctx context.Context) ([]*models.Post, error)
	UpdatePost(ctx context.Context, post *models.Post) error
	DeletePost(ctx context.Context, id uint) error
	GetPostsByUserId(ctx context.Context, userID uint) ([]*models.Post, error)
}

type postService struct {
	repo repository.PostRepository
}

func NewPostService(repo repository.PostRepository) PostService {
	return &postService{repo: repo}
}

func (s *postService) CreatePost(ctx context.Context, post *models.Post) error {
	if err := validatePost(post); err != nil {
		return err
	}
	post.UUID = uuid.NewString()
	return s.repo.CreatePost(ctx, post)
}
func (s *postService) GetPost(ctx context.Context, id uint) (*models.Post, error) {
	return s.repo.GetPost(ctx, id)
}
func (s *postService) GetAllPosts(ctx context.Context) ([]*models.Post, error) {
	return s.repo.GetAllPosts(ctx)
}
func (s *postService) UpdatePost(ctx context.Context, post *models.Post) error {
	if post.ID == 0 {
		return errors.New("id should be positive")
	}
	return s.repo.UpdatePost(ctx, post)
}
func (s *postService) DeletePost(ctx context.Context, id uint) error {
	return s.repo.DeletePost(ctx, id)
}
func (s *postService) GetPostsByUserId(ctx context.Context, userID uint) ([]*models.Post, error) {
	return s.repo.GetPostsByUserId(ctx, userID)
}

func validatePost(post *models.Post) error {
	if post.Content == "" {
		return errors.New("content required")
	}
	if post.UserID <= 0 {
		return errors.New("User ID should be positive")
	}
	return nil
}
