package valueobjects

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrInvalidUUID = errors.New("invalid uuid")
)

type EntityID struct {
	uuid uuid.UUID
}

func NewEntityID() EntityID {
	return EntityID{uuid: uuid.New()}
}

func NewEntityIDFromString(id string) (EntityID, error) {
	uuid, err := uuid.Parse(id)

	if err != nil {
		return EntityID{}, ErrInvalidUUID
	}

	return EntityID{uuid: uuid}, nil
}

func (e EntityID) String() string {
	return e.uuid.String()
}

func (e EntityID) Equals(other EntityID) bool {
	return e.uuid == other.uuid
}

func (e EntityID) IsEmpty() bool {
	return e.uuid == uuid.Nil
}
