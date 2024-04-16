package events

import (
	"time"

	valueobjects "github.com/harlancleiton/go-tweets/pkg/domain/value_objects"
)

type TweetCreatedEventPayload struct {
	Text      string
	AuthorID  valueobjects.EntityID
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TweetCreatedEvent struct {
	id        valueobjects.EntityID
	entityId  valueobjects.EntityID
	ocurredAt time.Time
	payload   TweetCreatedEventPayload
}

func NewCreatedTweetEvent(entityId *valueobjects.EntityID, text string, authorID valueobjects.EntityID, tweetCreatedAt, tweetUpdatedAt time.Time) (*TweetCreatedEvent, error) {
	id, err := valueobjects.NewEntityID()

	if err != nil {
		return nil, err
	}

	return &TweetCreatedEvent{
		id:        *id,
		entityId:  *entityId,
		ocurredAt: time.Now(),
		payload: TweetCreatedEventPayload{
			Text:      text,
			AuthorID:  authorID,
			CreatedAt: tweetCreatedAt,
			UpdatedAt: tweetUpdatedAt,
		},
	}, nil
}

func FromExistingCreatedTweetEvent(id, entityId, text, authorId string, ocurredAt time.Time) (*TweetCreatedEvent, error) {
	eventId, err := valueobjects.NewEntityIDFromString(id)

	if err != nil {
		return nil, err
	}

	eId, err := valueobjects.NewEntityIDFromString(entityId)

	if err != nil {
		return nil, err
	}

	a, err := valueobjects.NewEntityIDFromString(authorId)

	if err != nil {
		return nil, err
	}

	return &TweetCreatedEvent{
		id:        *eventId,
		entityId:  *eId,
		ocurredAt: ocurredAt,
		payload: TweetCreatedEventPayload{
			Text:     text,
			AuthorID: *a,
		},
	}, nil
}

func (c *TweetCreatedEvent) ID() valueobjects.EntityID {
	return c.id
}

func (c *TweetCreatedEvent) EntityID() valueobjects.EntityID {
	return c.entityId
}

func (c *TweetCreatedEvent) Name() string {
	return "CreatedEvent"
}

func (c *TweetCreatedEvent) OcurredAt() time.Time {
	return c.ocurredAt
}

func (c *TweetCreatedEvent) Payload() interface{} {
	return c.payload
}

func (c *TweetCreatedEvent) EventPayload() TweetCreatedEventPayload {
	return c.payload
}
