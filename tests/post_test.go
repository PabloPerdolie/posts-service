package tests

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"graphql-comments/internal/models"
	mock_storage "graphql-comments/internal/storage/mocks"
	"testing"
)

func TestAddPost(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_storage.NewMockPostRepository(ctrl)
	newPost := &models.Post{Title: "Test Post", Content: "Test Content"}

	mockRepo.EXPECT().AddPost(newPost).Return(newPost, nil).Times(1)

	result, err := mockRepo.AddPost(newPost)
	assert.NoError(t, err)
	assert.Equal(t, newPost, result)
}

func TestGetPost(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_storage.NewMockPostRepository(ctrl)
	postID := "123"
	expectedPost := &models.Post{ID: postID, Title: "Test Post", Content: "Test Content"}

	mockRepo.EXPECT().GetPost(postID).Return(expectedPost, nil).Times(1)

	result, err := mockRepo.GetPost(postID)
	assert.NoError(t, err)
	assert.Equal(t, expectedPost, result)
}

func TestGetPostNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_storage.NewMockPostRepository(ctrl)
	postID := "123"

	mockRepo.EXPECT().GetPost(postID).Return(nil, errors.New("post not found")).Times(1)

	result, err := mockRepo.GetPost(postID)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "post not found", err.Error())
}

func TestGetPosts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_storage.NewMockPostRepository(ctrl)
	expectedPosts := []*models.Post{
		{ID: "1", Title: "Post 1", Content: "Content 1"},
		{ID: "2", Title: "Post 2", Content: "Content 2"},
	}

	mockRepo.EXPECT().GetPosts().Return(expectedPosts, nil).Times(1)

	result, err := mockRepo.GetPosts()
	assert.NoError(t, err)
	assert.Equal(t, expectedPosts, result)
}
