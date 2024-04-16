package dto

import "github.com/harlancleiton/go-tweets/internal/domain/entities"

type AuthorDto struct {
	ID       string
	Username string
}

func NewAuthorDto(author *entities.Author) *AuthorDto {
	aId := author.ID()

	return &AuthorDto{
		ID:       aId.String(),
		Username: author.Username(),
	}
}
