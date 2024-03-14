package entities

import (
	"testing"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
)

func TestNewTweet(t *testing.T) {
	author, err := NewUserFromExisting(faker.UUIDHyphenated(), faker.Username(), time.Now(), time.Now(), true, true)

	assert.NoError(t, err)

	text := faker.Paragraph()

	tweet, err := NewTweet(text, *author)

	assert.NoError(t, err)
	assert.Equal(t, text, tweet.Text())
	assert.Equal(t, *author, tweet.Author())
}

func TestTweetText(t *testing.T) {
	author, err := NewUserFromExisting(faker.UUIDHyphenated(), faker.Username(), time.Now(), time.Now(), true, true)

	assert.NoError(t, err)

	text := faker.Paragraph()

	tweet, _ := NewTweet(text, *author)

	assert.Equal(t, text, tweet.Text())
}

func TestTweetAuthor(t *testing.T) {
	author, err := NewUserFromExisting(faker.UUIDHyphenated(), faker.Username(), time.Now(), time.Now(), true, true)

	assert.NoError(t, err)

	text := faker.Paragraph()

	tweet, _ := NewTweet(text, *author)

	assert.Equal(t, *author, tweet.Author())
}
