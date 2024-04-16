package services

import (
	"context"
	"time"

	"github.com/harlancleiton/go-tweets/internal/domain/dto"
	"github.com/harlancleiton/go-tweets/internal/domain/entities"
	"github.com/harlancleiton/go-tweets/internal/domain/repositories"
	"github.com/harlancleiton/go-tweets/internal/domain/services"
	"github.com/harlancleiton/go-tweets/pkg/domain/events"
)

type TweetService struct {
	tweetService    *services.TweetService
	userRepository  repositories.UserRepository
	tweetRepository repositories.TweetRepository
	dispatcher      events.EventDispatcher
}

func (s *TweetService) CreateTweet(username string, input *dto.CreateTweetInput) (*dto.TweetDto, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	author, err := s.userRepository.FindByUsername(ctx, username)

	if err != nil {
		return nil, err
	}

	f := entities.NewTweetFactory(s.dispatcher)
	t, err := f.CreateNewTweet(input.Text, author)

	if err != nil {
		return nil, err
	}

	err = s.tweetService.Tweet(author, t)

	if err != nil {
		return nil, err
	}

	err = s.tweetRepository.Save(ctx, t)

	if err != nil {
		return nil, err
	}

	t.Commit()

	return dto.NewTweetDto(t), nil
}

func NewTweetService(userRepository repositories.UserRepository, tweetRepository repositories.TweetRepository, dispatcher events.EventDispatcher) *TweetService {
	return &TweetService{
		tweetService:    services.NewTweetService(),
		tweetRepository: tweetRepository,
		userRepository:  userRepository,
		dispatcher:      dispatcher,
	}
}
