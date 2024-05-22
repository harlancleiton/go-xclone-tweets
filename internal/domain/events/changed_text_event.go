package events

import (
	"time"

	valueobjects "github.com/harlancleiton/go-tweets/pkg/domain/value_objects"
)

type TweetChangedTextEventPayload struct {
	OldText string `json:"old_text"`
	NewText string `json:"new_text"`
}

type TweetChangedTextEvent struct {
	id        valueobjects.EntityID
	entityId  valueobjects.EntityID
	ocurredAt time.Time
	payload   TweetChangedTextEventPayload
}

func (c *TweetChangedTextEvent) ID() valueobjects.EntityID {
	return c.id
}

func (c *TweetChangedTextEvent) EntityID() valueobjects.EntityID {
	return c.entityId
}

func (c *TweetChangedTextEvent) EntityName() string {
	return "Tweet"
}

func (c *TweetChangedTextEvent) Name() string {
	return "TweetCreated"
}

func (c *TweetChangedTextEvent) OcurredAt() time.Time {
	return c.ocurredAt
}

func (c *TweetChangedTextEvent) Payload() interface{} {
	return c.payload
}

func (c *TweetChangedTextEvent) EventPayload() TweetChangedTextEventPayload {
	return c.payload
}

func NewChangedTweetTextEvent(entityId valueobjects.EntityID, oldText, newText string) (*TweetChangedTextEvent, error) {
	id, err := valueobjects.NewEntityID()

	if err != nil {
		return nil, err
	}

	return &TweetChangedTextEvent{
		id:        *id,
		entityId:  entityId,
		ocurredAt: time.Now(),
		payload: TweetChangedTextEventPayload{
			OldText: oldText,
			NewText: newText,
		},
	}, nil
}
