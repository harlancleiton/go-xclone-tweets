package memory

import (
	"context"
	"time"

	"github.com/harlancleiton/go-tweets/internal/domain/entities"
)

type MemoryTweetRepository struct {
	tweets []*entities.Tweet
}

func (r *MemoryTweetRepository) Save(ctx context.Context, tweet *entities.Tweet) error {
	select {
	case <-time.After(200 * time.Millisecond):
		r.tweets = append(r.tweets, tweet)
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func NewMemoryTweetRepository() *MemoryTweetRepository {
	return &MemoryTweetRepository{}
}
