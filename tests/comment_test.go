package tests

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"graphql-comments/internal/models"
	mock_storage "graphql-comments/internal/storage/mocks"
	"testing"
)

func TestAddComment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_storage.NewMockCommentRepository(ctrl)
	newComment := &models.Comment{Content: "Test Comment"}

	mockRepo.EXPECT().AddComment(newComment).Return(newComment, nil).Times(1)

	result, err := mockRepo.AddComment(newComment)
	assert.NoError(t, err)
	assert.Equal(t, newComment, result)
}

func TestAddCommentError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_storage.NewMockCommentRepository(ctrl)
	newComment := &models.Comment{Content: "Test Comment"}

	mockRepo.EXPECT().AddComment(newComment).Return(nil, errors.New("failed to add comment")).Times(1)

	result, err := mockRepo.AddComment(newComment)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "failed to add comment", err.Error())
}

func TestGetComments(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_storage.NewMockCommentRepository(ctrl)
	postID := "123"
	expectedComments := []*models.Comment{
		{Content: "Test Comment 1"},
		{Content: "Test Comment 2"},
	}

	mockRepo.EXPECT().GetComments(postID, 10, 0).Return(expectedComments, nil).Times(1)

	result, err := mockRepo.GetComments(postID, 10, 0)
	assert.NoError(t, err)
	assert.Equal(t, expectedComments, result)
}

func TestGetCommentsWithPagination(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_storage.NewMockCommentRepository(ctrl)
	postID := "123"
	limit := 5
	offset := 10
	expectedComments := []*models.Comment{
		{Content: "Comment 11"},
		{Content: "Comment 12"},
	}

	mockRepo.EXPECT().GetComments(postID, limit, offset).Return(expectedComments, nil).Times(1)

	result, err := mockRepo.GetComments(postID, limit, offset)
	assert.NoError(t, err)
	assert.Equal(t, expectedComments, result)
}

func TestGetCommentsError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_storage.NewMockCommentRepository(ctrl)
	postID := "123"

	mockRepo.EXPECT().GetComments(postID, 10, 0).Return(nil, errors.New("failed to get comments")).Times(1)

	result, err := mockRepo.GetComments(postID, 10, 0)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "failed to get comments", err.Error())
}

func TestManageComments(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_storage.NewMockPostRepository(ctrl)
	postID := "1"
	enable := true
	expectedPost := &models.Post{ID: postID, CommentsEnabled: enable}

	mockRepo.EXPECT().ManageComments(postID, enable).Return(expectedPost, nil).Times(1)

	result, err := mockRepo.ManageComments(postID, enable)
	assert.NoError(t, err)
	assert.Equal(t, expectedPost, result)
}
