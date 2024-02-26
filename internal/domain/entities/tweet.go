package entities

import (
	"time"

	valueobjects "github.com/harlancleiton/go-tweets/pkg/domain/value_objects"
)

type Tweet struct {
	ID        valueobjects.EntityID
	Text      string
	CreatedAt time.Time
}

func NewTweet(text string) (*Tweet, error) {
	return &Tweet{
		ID:        valueobjects.NewEntityID(),
		Text:      text,
		CreatedAt: time.Now(),
	}, nil
}
