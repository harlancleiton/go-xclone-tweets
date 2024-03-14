package services

import (
	"errors"

	"github.com/harlancleiton/go-tweets/internal/domain/entities"
)

type TweetService struct{}

var (
	ErrAuthorCannotTweet = errors.New("author cannot tweet")
)

func (s *TweetService) Tweet(author *entities.Author, tweet *entities.Tweet) error {
	if !author.CanTweet() {
		return ErrAuthorCannotTweet
	}

	author.AddTweet(tweet)

	return nil
}

func NewTweetService() *TweetService {
	return &TweetService{}
}
