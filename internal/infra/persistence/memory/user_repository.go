package memory

import (
	"context"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/harlancleiton/go-tweets/internal/domain/entities"
)

type MemoryUserRepository struct {
	users []*entities.Author
}

func (r *MemoryUserRepository) FindByUsername(ctx context.Context, username string) (*entities.Author, error) {
	for _, user := range r.users {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			if user.Username() == username {
				return user, nil
			}
		}
	}

	return nil, nil
}

func NewMemoryUserRepository() *MemoryUserRepository {
	user, err := entities.NewUserFromExisting(faker.UUIDHyphenated(), "somebody", time.Now(), time.Now(), false, true)

	if err != nil {
		panic(err)
	}

	return &MemoryUserRepository{
		users: []*entities.Author{
			user,
		},
	}
}
