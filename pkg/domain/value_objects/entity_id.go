package valueobjects

import (
	"errors"

	"github.com/google/uuid"
)

// ErrInvalidUUID is an error thrown when the provided UUID string is invalid.
var ErrInvalidUUID = errors.New("invalid uuid")

// EntityID is a value object that represents a unique identifier for entities.
type EntityID struct {
	uuid uuid.UUID // UUID is a Universally Unique Identifier
}

// NewEntityID generates a new EntityID with a random UUID.
func NewEntityID() EntityID {
	return EntityID{uuid: uuid.New()}
}

// NewEntityIDFromString creates a new EntityID from a UUID given as a string.
// Returns an error if the provided string is not a valid UUID.
//
// Example:
//
//	id, err := NewEntityIDFromString("d9b2d63d-1c4a-4b89-8ba4-b7114aecc446")
//	if err != nil {
//	    log.Fatalf("failed to create entity id: %v", err)
//	}
func NewEntityIDFromString(id string) (EntityID, error) {
	uuid, err := uuid.Parse(id)

	if err != nil {
		return EntityID{}, ErrInvalidUUID
	}

	return EntityID{uuid: uuid}, nil
}

// String returns string representation of the UUID
func (e EntityID) String() string {
	return e.uuid.String()
}

// Equals checks if the UUID of another EntityID_equal to the UUID of the current EntityID.
func (e EntityID) Equals(other EntityID) bool {
	return e.uuid == other.uuid
}

// IsEmpty checks if the UUID is Nil (equivalent to an empty UUID).
func (e EntityID) IsEmpty() bool {
	return e.uuid == uuid.Nil
}
