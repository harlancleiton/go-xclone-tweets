package repositories

import (
	"context"

	"github.com/harlancleiton/go-tweets/internal/domain/entities"
)

type TweetRepository interface {
	Save(ctx context.Context, tweet *entities.Tweet) error
}
