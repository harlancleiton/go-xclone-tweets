package entities

import (
	"errors"

	tweetEvents "github.com/harlancleiton/go-tweets/internal/domain/events"

	eventsPkg "github.com/harlancleiton/go-tweets/pkg/domain/events"
	valueobjects "github.com/harlancleiton/go-tweets/pkg/domain/value_objects"
)

var (
	ErrTextCannotBeEmpty              = errors.New("text cannot be empty")
	ErrAuthorCannotBeNil              = errors.New("author cannot be nil")
	ErrNoEvents                       = errors.New("no events provided")
	ErrFirstEventNotCreatedTweetEvent = errors.New("the first event is not a CreatedTweetEvent")
)

type TweetFactory struct {
	dispatcher eventsPkg.EventDispatcher
}

func NewTweetFactory(dispatcher eventsPkg.EventDispatcher) *TweetFactory {
	return &TweetFactory{
		dispatcher: dispatcher,
	}
}

func (f *TweetFactory) CreateNewTweet(text string, author *Author) (*Tweet, error) {
	if text == "" {
		return nil, ErrTextCannotBeEmpty
	}

	if author == nil {
		return nil, ErrAuthorCannotBeNil
	}

	entityId, err := valueobjects.NewEntityID()

	if err != nil {
		return nil, err
	}

	return NewTweet(entityId, text, author, f.dispatcher)
}

func (f *TweetFactory) CreateTweetFromEvents(events []eventsPkg.Event, author *Author) (*Tweet, error) {
	if len(events) == 0 {
		return nil, ErrNoEvents
	}

	firstEvent := events[0]
	createdEvent, ok := firstEvent.(*tweetEvents.TweetCreatedEvent)

	if !ok {
		return nil, ErrFirstEventNotCreatedTweetEvent
	}

	aId := createdEvent.EventPayload().AuthorID
	a, err := NewUserFromExisting(aId.String(), author.Username(), author.CreatedAt(), author.UpdatedAt(), author.IsBlocked(), author.EmailVerified())

	if err != nil {
		return nil, err
	}

	t, err := NewTweetFromEvents(createdEvent, events, a, f.dispatcher)

	if err != nil {
		return nil, err
	}

	return t, nil
}
