package services

import (
	"context"
	"time"

	"github.com/harlancleiton/go-tweets/internal/domain/repositories"
)

type UserStatusService struct {
	userRepository repositories.UserRepository
}

func (s *UserStatusService) CanPostTweet(username string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	author, err := s.userRepository.FindByUsername(ctx, username)

	if err != nil {
		return false, err
	}

	return !author.IsBlocked() || author.EmailVerified(), nil
}

func NewUserStatusService(userRepository repositories.UserRepository) *UserStatusService {
	return &UserStatusService{userRepository: userRepository}
}
