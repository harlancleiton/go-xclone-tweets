package valueobjects

import (
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewEntityID(t *testing.T) {
	eid := NewEntityID()

	assert.False(t, eid.IsEmpty(), "NewEntityID() should not return an empty UUID")
}

func TestNewEntityIDFromString_Valid(t *testing.T) {
	validUUID := faker.UUIDHyphenated()

	eid, err := NewEntityIDFromString(validUUID)
	assert.NoError(t, err, "Unexpected error")

	assert.Equal(t, validUUID, eid.String(), "Expected UUID does not match the given UUID")
}

func TestNewEntityIDFromString_Invalid(t *testing.T) {
	invalidUUID := "invalid"

	_, err := NewEntityIDFromString(invalidUUID)
	assert.Error(t, err, "Expected error did not occur")

	assert.Equal(t, ErrInvalidUUID, err, "Expected ErrInvalidUUID, but got another error")
}

func TestEntityID_String(t *testing.T) {
	validUUID := faker.UUIDHyphenated()
	eid, _ := NewEntityIDFromString(validUUID)

	assert.Equal(t, validUUID, eid.String(), "Expected UUID does not match the given UUID")
}

func TestEntityID_Equals(t *testing.T) {
	validUUID := faker.UUIDHyphenated()
	eid1, _ := NewEntityIDFromString(validUUID)
	eid2, _ := NewEntityIDFromString(validUUID)

	assert.True(t, eid1.Equals(eid2), "Expected EntityIDs to be equal")
}

func TestEntityID_IsEmpty(t *testing.T) {
	emptyUUID := uuid.Nil
	eid := EntityID{uuid: emptyUUID}

	assert.True(t, eid.IsEmpty(), "Expected UUID to be empty")
}
