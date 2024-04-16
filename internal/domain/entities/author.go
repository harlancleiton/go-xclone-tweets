package entities

import (
	"time"

	valueobjects "github.com/harlancleiton/go-tweets/pkg/domain/value_objects"
)

type Author struct {
	BaseEntity
	username      string
	isBlocked     bool
	emailVerified bool
	tweets        []*Tweet
}

func (u *Author) Username() string {
	return u.username
}

func (u *Author) IsBlocked() bool {
	return u.isBlocked
}

func (u *Author) EmailVerified() bool {
	return u.emailVerified
}

func (u *Author) CanTweet() bool {
	return !u.isBlocked && u.emailVerified
}

func (u *Author) AddTweet(tweet *Tweet) {
	u.tweets = append(u.tweets, tweet)
}

func NewUserFromExisting(id, username string, createdAt, updatedAt time.Time, isBlocked, emailVerified bool) (*Author, error) {
	entityId, err := valueobjects.NewEntityIDFromString(id)

	if err != nil {
		return nil, err
	}

	// TODO add validation
	return &Author{
		BaseEntity: BaseEntity{
			id:        *entityId,
			createdAt: createdAt,
			updatedAt: updatedAt,
		},
		username:      username,
		isBlocked:     isBlocked,
		emailVerified: emailVerified,
	}, nil
}
