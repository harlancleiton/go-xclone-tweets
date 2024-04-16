package entities

import (
	"errors"
	"time"

	"github.com/harlancleiton/go-tweets/pkg/domain/events"
	valueobjects "github.com/harlancleiton/go-tweets/pkg/domain/value_objects"
)

var (
	ErrUnknownEvent = errors.New("unknown event")
)

type BaseEntity struct {
	id         valueobjects.EntityID // The unique identifier for the Entity
	createdAt  time.Time             // The timestamp the Entity was created
	updatedAt  time.Time             // The timestamp the Entity was last updated
	dispatcher events.EventDispatcher
	events     []events.Event
	version    uint32
}

func (b *BaseEntity) Equals(other *BaseEntity) bool {
	return b.id.Equals(other.id)
}

func (b *BaseEntity) ID() valueobjects.EntityID {
	return b.id
}

func (b *BaseEntity) CreatedAt() time.Time {
	return b.createdAt
}

func (b *BaseEntity) UpdatedAt() time.Time {
	return b.updatedAt
}

func (b *BaseEntity) Commit() error {
	for _, event := range b.events {
		err := b.dispatcher.Dispatch(event)

		if err != nil {
			return err
		}
	}

	b.version += uint32(len(b.events))
	b.events = []events.Event{}
	return nil
}

func (b *BaseEntity) Uncommit() {
	b.events = []events.Event{}
}
