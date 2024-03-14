package entities

import (
	"time"

	valueobjects "github.com/harlancleiton/go-tweets/pkg/domain/value_objects"
)

type BaseEntity struct {
	id        valueobjects.EntityID // The unique identifier for the Entity
	createdAt time.Time             // The timestamp the Entity was created
	updatedAt time.Time             // The timestamp the Entity was last updated
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
