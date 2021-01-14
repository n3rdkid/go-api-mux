package service

import (
	"errors"
	"go-api-mux/entity"
	"go-api-mux/repository"
	"math/rand"
)

// PostService ->
type PostService interface {
	Validate(post *entity.Post) error
	Create(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}

type service struct{}

var (
	repo repository.PostRepository
)

// NewPostService -> Constructor function
func NewPostService(repository repository.PostRepository) PostService {
	repo = repository
	return &service{}
}

func (*service) Validate(post *entity.Post) error {
	if post == nil {
		err := errors.New("the post is empty")
		return err
	}
	if post.Title == "" {
		err := errors.New("the title field is required")
		return err
	}

	return nil

}
func (*service) Create(post *entity.Post) (*entity.Post, error) {
	post.ID = rand.Int63()
	return repo.Save(post)
}
func (*service) FindAll() ([]entity.Post, error) {
	return repo.FindAll()
}
