package usecases

import (
	"context"

	"github.com/harlancleiton/go-tweets/internal/dto"
	"github.com/harlancleiton/go-tweets/internal/entities"
)

type CreateTweet struct{}

func (s *CreateTweet) Execute(ctx context.Context, input dto.CreateTweetInput) (*dto.TweetDto, error) {
	tweet, err := entities.NewTweet(input.Text)

	if err != nil {
		return nil, err
	}

	return &dto.TweetDto{
		ID:        tweet.ID.String(),
		Text:      tweet.Text,
		CreatedAt: tweet.CreatedAt,
	}, nil
}

func NewCreateTweet() *CreateTweet {
	return &CreateTweet{}
}
