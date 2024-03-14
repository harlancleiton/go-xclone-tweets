package valueobjects

import (
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewEntityID(t *testing.T) {
	id, err := NewEntityID()

	assert.NoError(t, err)
	assert.False(t, id.IsEmpty())
}

func TestNewEntityIDFromString_Valid(t *testing.T) {
	validUUID := faker.UUIDHyphenated()

	id, err := NewEntityIDFromString(validUUID)

	assert.NoError(t, err)
	assert.Equal(t, validUUID, id.String())
}

func TestNewEntityIDFromString_Invalid(t *testing.T) {
	invalidUUID := "invalid"

	_, err := NewEntityIDFromString(invalidUUID)
	assert.Error(t, err)

	assert.Equal(t, ErrInvalidUUID, err)
}

func TestEntityID_String(t *testing.T) {
	validUUID := faker.UUIDHyphenated()
	eid, _ := NewEntityIDFromString(validUUID)

	assert.Equal(t, validUUID, eid.String())
}

func TestEntityID_Equals(t *testing.T) {
	validUUID := faker.UUIDHyphenated()
	eid1, _ := NewEntityIDFromString(validUUID)
	eid2, _ := NewEntityIDFromString(validUUID)

	assert.True(t, eid1.Equals(eid2))
}

func TestEntityID_IsEmpty(t *testing.T) {
	emptyUUID := uuid.Nil
	eid := EntityID{uuid: emptyUUID}

	assert.True(t, eid.IsEmpty())
}
