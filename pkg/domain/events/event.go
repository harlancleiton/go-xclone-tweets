package events

import (
	"time"

	valueobjects "github.com/harlancleiton/go-tweets/pkg/domain/value_objects"
)

type Event interface {
	ID() valueobjects.EntityID
	EntityID() valueobjects.EntityID
	Name() string
	OcurredAt() time.Time
	Payload() interface{}
}
