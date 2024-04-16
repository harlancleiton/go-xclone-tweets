package dto

import (
	"time"

	"github.com/harlancleiton/go-tweets/internal/domain/entities"
)

type TweetDto struct {
	ID        string
	Text      string
	CreatedAt time.Time
	Author    AuthorDto
}

func NewTweetDto(tweet *entities.Tweet) *TweetDto {
	tId := tweet.ID()
	a := tweet.Author()

	return &TweetDto{
		ID:        tId.String(),
		Text:      tweet.Text(),
		CreatedAt: tweet.CreatedAt(),
		Author:    *NewAuthorDto(&a)}
}
