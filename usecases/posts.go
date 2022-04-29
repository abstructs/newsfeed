package usecases

import (
	"context"

	"github.com/abstructs/newsfeed/models"
	"go.uber.org/zap"
)

type IPostUsecase interface {
	CreatePost(ctx context.Context, postRequest models.PostRequest) (models.Post, error)
	GetPosts(ctx context.Context) []models.Post
}

type Usecase struct {
	logger    *zap.Logger
	posts     []models.Post
	publisher models.IPublisher
}

func NewUsecase(logger *zap.Logger, publisher models.IPublisher) IPostUsecase {
	return &Usecase{
		logger:    logger,
		posts:     []models.Post{},
		publisher: publisher,
	}
}

func (uc *Usecase) CreatePost(ctx context.Context, postRequest models.PostRequest) (models.Post, error) {
	newPost := models.Post{
		Text: postRequest.Text,
	}

	err := uc.publisher.Publish(ctx, postRequest.PostType, postRequest.Text)
	if err != nil {
		uc.logger.Sugar().Errorf("Failed to publish create post event:", err)
		return newPost, err
	}
	uc.logger.Sugar().Debugf("Successfullly publisher create post event")

	uc.posts = append(uc.posts, newPost)

	return newPost, nil
}

func (uc *Usecase) GetPosts(ctx context.Context) []models.Post {
	return uc.posts
}
