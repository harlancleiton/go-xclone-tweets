package services

import (
	"context"
	"time"

	"github.com/harlancleiton/go-tweets/internal/domain/dto"
	"github.com/harlancleiton/go-tweets/internal/domain/entities"
	"github.com/harlancleiton/go-tweets/internal/domain/repositories"
	"github.com/harlancleiton/go-tweets/internal/domain/services"
)

type TweetService struct {
	tweetService    *services.TweetService
	userRepository  repositories.UserRepository
	tweetRepository repositories.TweetRepository
}

func (s *TweetService) CreateTweet(username string, input *dto.CreateTweetInput) (*dto.TweetDto, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	author, err := s.userRepository.FindByUsername(ctx, username)

	if err != nil {
		return nil, err
	}

	tweet, err := entities.NewTweet(input.Text, *author)

	if err != nil {
		return nil, err
	}

	err = s.tweetService.Tweet(author, tweet)

	if err != nil {
		return nil, err
	}

	err = s.tweetRepository.Save(ctx, tweet)

	if err != nil {
		return nil, err
	}

	return dto.NewTweetDto(tweet), nil
}

func NewTweetService(userRepository repositories.UserRepository, tweetRepository repositories.TweetRepository) *TweetService {
	return &TweetService{
		tweetService:    services.NewTweetService(),
		tweetRepository: tweetRepository,
		userRepository:  userRepository,
	}
}
