package events

import (
	"time"

	valueobjects "github.com/harlancleiton/go-tweets/pkg/domain/value_objects"
)

type TweetCreatedEventPayload struct {
	Text     string `json:"text"`
	AuthorID string `json:"author_id"`
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
			Text:     text,
			AuthorID: authorID.String(),
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
			AuthorID: a.String(),
		},
	}, nil
}

func (c *TweetCreatedEvent) ID() valueobjects.EntityID {
	return c.id
}

func (c *TweetCreatedEvent) EntityID() valueobjects.EntityID {
	return c.entityId
}

func (c *TweetCreatedEvent) EntityName() string {
	return "Tweet"
}

func (c *TweetCreatedEvent) Name() string {
	return "TweetCreated"
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
