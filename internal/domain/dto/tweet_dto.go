package dto

import (
	"time"

	"github.com/harlancleiton/go-tweets/internal/domain/entities"
)

type AuthorDto struct {
	ID       string
	Username string
}

type TweetDto struct {
	ID        string
	Text      string
	CreatedAt time.Time
	Author    AuthorDto
}

func NewTweetDto(tweet *entities.Tweet) *TweetDto {
	author := tweet.Author()

	return &TweetDto{
		ID:        tweet.ID().String(),
		Text:      tweet.Text(),
		CreatedAt: tweet.CreatedAt(),
		Author: AuthorDto{
			ID:       author.ID().String(),
			Username: author.Username(),
		}}
}
