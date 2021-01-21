package service

import (
	"go-api-mux/entity"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (mock MockRepository) Save(post *entity.Post) (*entity.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*entity.Post), args.Error(1)
}
func (mock MockRepository) FindAll() ([]entity.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entity.Post), args.Error(1)
}
func TestValidateEmptyPost(t *testing.T) {
	testService := NewPostService(nil)
	err := testService.Validate(nil)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "the post is empty")

}

func TestValidateEmptyTitle(t *testing.T) {
	post := entity.Post{ID: 1, Title: "", Text: ""}
	testService := NewPostService(nil)
	err := testService.Validate(&post)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "the title field is required")
}

func TestFindAll(t *testing.T) {
	mockRepo := new(MockRepository)
	var identifier int64 = 1
	post := entity.Post{ID: identifier, Title: "A", Text: "B"}
	mockRepo.On("FindAll").Return([]entity.Post{post}, nil)
	testService := NewPostService((mockRepo))
	result, _ := testService.FindAll()
	mockRepo.AssertExpectations(t)
	assert.Equal(t, identifier, result[0].ID)
	assert.Equal(t, "A", result[0].Title)
	assert.Equal(t, "B", result[0].Text)

}
func TestCreate(t *testing.T) {
	mockRepo := new(MockRepository)
	post := entity.Post{Title: "A", Text: "B"}
	mockRepo.On("Save").Return(&post, nil)
	testService := NewPostService(mockRepo)
	result, err := testService.Create(&post)
	mockRepo.AssertExpectations(t)
	assert.NotNil(t, result.ID)
	assert.Equal(t, "A", result.Title)
	assert.Equal(t, "B", result.Text)
	assert.Nil(t, err)
}
