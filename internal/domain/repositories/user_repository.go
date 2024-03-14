package repositories

import (
	"context"

	"github.com/harlancleiton/go-tweets/internal/domain/entities"
)

type UserRepository interface {
	FindByUsername(ctx context.Context, username string) (*entities.Author, error)
}
