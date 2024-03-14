package infra

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
	case <-time.After(1 * time.Second):
		r.tweets = append(r.tweets, tweet)
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func NewMemoryTweetRepository() *MemoryTweetRepository {
	return &MemoryTweetRepository{}
}
