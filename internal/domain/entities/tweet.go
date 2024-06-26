package entities

import (
	"time"

	tweetEvents "github.com/harlancleiton/go-tweets/internal/domain/events"
	events "github.com/harlancleiton/go-tweets/pkg/domain/events"
	valueobjects "github.com/harlancleiton/go-tweets/pkg/domain/value_objects"
)

type Tweet struct {
	BaseEntity
	text   string
	author Author
}

func NewTweet(id *valueobjects.EntityID, text string, author *Author, dispatcher events.EventDispatcher) (*Tweet, error) {
	t := &Tweet{
		BaseEntity: BaseEntity{
			id:         *id,
			createdAt:  time.Now(),
			updatedAt:  time.Now(),
			dispatcher: dispatcher,
			version:    1,
			events:     []events.Event{},
		},
		text:   text,
		author: *author,
	}

	e, err := tweetEvents.NewCreatedTweetEvent(id, text, author.ID(), t.CreatedAt(), t.UpdatedAt())

	if err != nil {
		return nil, err
	}

	t.apply(e)

	return t, nil
}

func NewTweetFromEvents(createdEvent *tweetEvents.TweetCreatedEvent, remainingEvents []events.Event, author *Author, dispatcher events.EventDispatcher) (*Tweet, error) {
	t := &Tweet{
		BaseEntity: BaseEntity{
			id:         createdEvent.EntityID(),
			createdAt:  createdEvent.OcurredAt(),
			updatedAt:  createdEvent.OcurredAt(),
			dispatcher: dispatcher,
			version:    1,
			events:     []events.Event{},
		},
		text:   createdEvent.EventPayload().Text,
		author: *author,
	}

	for _, event := range remainingEvents {
		if err := t.apply(event); err != nil {
			return nil, err
		}
	}

	t.events = []events.Event{}

	return t, nil
}

func (t *Tweet) Text() string {
	return t.text
}

func (t *Tweet) Author() Author {
	return t.author
}

func (t *Tweet) ChangeText(text string) error {
	e, err := tweetEvents.NewChangedTweetTextEvent(t.ID(), t.text, text)

	if err != nil {
		return err
	}

	return t.apply(e)
}

func (t *Tweet) apply(event events.Event) error {
	t.events = append(t.events, event)

	switch e := event.(type) {
	case *tweetEvents.TweetCreatedEvent:
		t.text = e.EventPayload().Text
	case *tweetEvents.TweetChangedTextEvent:
		t.text = e.EventPayload().NewText
		t.updatedAt = e.OcurredAt()
	default:
		return ErrUnknownEvent
	}

	return ErrUnknownEvent
}
