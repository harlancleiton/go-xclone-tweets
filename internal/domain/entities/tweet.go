package entities

import (
	"time"

	valueobjects "github.com/harlancleiton/go-tweets/pkg/domain/value_objects"
)

type Tweet struct {
	BaseEntity
	text   string
	author Author
}

func NewTweet(text string, author Author) (*Tweet, error) {
	// TODO add validation
	entityId, err := valueobjects.NewEntityID()

	if err != nil {
		return nil, err
	}

	return &Tweet{
		BaseEntity: BaseEntity{
			id:        *entityId,
			createdAt: time.Now(),
			updatedAt: time.Now(),
		},
		text:   text,
		author: author,
	}, nil
}

func (t *Tweet) Text() string {
	return t.text
}

func (t *Tweet) Author() Author {
	return t.author
}
