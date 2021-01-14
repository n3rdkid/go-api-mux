package repository

import (
	"go-api-mux/entity"
)

// PostRepository ->
type PostRepository interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}
