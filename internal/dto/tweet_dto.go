package dto

import (
	"time"
)

type TweetDto struct {
	ID        string
	Text      string
	CreatedAt time.Time
}
