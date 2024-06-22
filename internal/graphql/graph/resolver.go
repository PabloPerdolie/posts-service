package graph

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"
	"graphql-comments/internal/graphql/graph/generated"
	"graphql-comments/internal/graphql/graph/model"
	"graphql-comments/internal/models"
	"graphql-comments/internal/services"
	"log"
)

type Resolver struct {
	postService    services.PostService
	commentService services.CommentService
}

func NewResolver(postService services.PostService, commentService services.CommentService) *Resolver {
	return &Resolver{
		postService:    postService,
		commentService: commentService,
	}
}

func (r *mutationResolver) CreatePost(ctx context.Context, title string, content string, commentsEnabled bool) (*model.Post, error) {
	post, err := r.postService.CreatePost(title, content, commentsEnabled)
	return &model.Post{
		ID:              post.ID,
		Title:           post.Title,
		Content:         post.Content,
		CommentsEnabled: post.CommentsEnabled,
	}, err
}

func (r *mutationResolver) CreateComment(ctx context.Context, postID string, parentID *string, content string) (*model.Comment, error) {
	comment, err := r.commentService.CreateComment(postID, parentID, content)
	if err != nil {
		return nil, err
	}

	return convertToGQLComment(comment), nil
}

func (r *mutationResolver) ManageComments(ctx context.Context, postID string, enable bool) (*model.Post, error) {
	post, err := r.postService.ManageComments(postID, enable)
	if err != nil {
		return nil, err
	}

	return &model.Post{
		ID:              post.ID,
		Title:           post.Title,
		Content:         post.Content,
		CommentsEnabled: post.CommentsEnabled,
	}, nil
}

func (r *queryResolver) Posts(ctx context.Context) ([]*model.Post, error) {
	posts, err := r.postService.GetPosts()
	if err != nil {
		return nil, err
	}

	gqlPosts := make([]*model.Post, len(posts))
	for i, post := range posts {
		gqlPosts[i] = &model.Post{
			ID:              post.ID,
			Title:           post.Title,
			Content:         post.Content,
			CommentsEnabled: post.CommentsEnabled,
		}
	}

	return gqlPosts, nil
}

func (r *queryResolver) Post(ctx context.Context, id string) (*model.Post, error) {
	post, err := r.postService.GetPostByID(id)
	if err != nil {
		return nil, err
	}

	return &model.Post{
		ID:              post.ID,
		Title:           post.Title,
		Content:         post.Content,
		CommentsEnabled: post.CommentsEnabled,
	}, nil
}

func (r *queryResolver) Comments(ctx context.Context, postID string) ([]*model.Comment, error) {
	comments, err := r.commentService.GetCommentsByPostID(postID)
	if err != nil {
		return nil, err
	}

	gqlComments := make([]*model.Comment, len(comments))
	for i, comment := range comments {
		gqlComments[i] = convertToGQLComment(comment)
	}

	return gqlComments, nil
}

func (r *subscriptionResolver) CommentAdded(ctx context.Context, postID string) (<-chan *model.Comment, error) {
	commentChan := make(chan *model.Comment)

	comChan, err := r.commentService.SubscribeToComments(postID)
	if err != nil {
		return nil, err
	}

	go func() {
		defer close(commentChan) // Закройте канал при завершении подписки

		for {
			select {
			case <-ctx.Done():
				log.Println("Subscription ended")
				return
			case comment := <-comChan:
				commentChan <- &model.Comment{
					ID:        comment.ID,
					PostID:    comment.PostID,
					ParentID:  comment.ParentID,
					Content:   comment.Content,
					CreatedAt: comment.CreatedAt.String(),
				}
			}
		}
	}()

	return commentChan, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }

func convertToGQLComment(comment *models.Comment) *model.Comment {
	gqlChildren := make([]*model.Comment, len(comment.Children))
	for i, child := range comment.Children {
		gqlChildren[i] = convertToGQLComment(child)
	}

	return &model.Comment{
		ID:        comment.ID,
		PostID:    comment.PostID,
		ParentID:  comment.ParentID,
		Content:   comment.Content,
		CreatedAt: comment.CreatedAt.String(),
		Children:  gqlChildren,
	}
}
