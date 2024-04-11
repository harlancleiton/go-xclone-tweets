package grpc

import (
	"context"

	"github.com/harlancleiton/go-tweets/internal/application/services"
	"github.com/harlancleiton/go-tweets/internal/domain/dto"
	"github.com/harlancleiton/go-tweets/internal/domain/entities"
	grpcInterceptor "github.com/harlancleiton/go-tweets/pkg/middlewares/grpc"
	"github.com/harlancleiton/go-tweets/pkg/pb"
)

type GrpcCreateTweetHandler struct {
	tweetService *services.TweetService
}

func (h *GrpcCreateTweetHandler) Create(ctx context.Context, request *pb.CreateTweetRequest) (*pb.TweetResponse, error) {
	user, ok := ctx.Value(grpcInterceptor.AuthenticatedUser).(*entities.Author)

	if !ok {
		return nil, grpcInterceptor.ErrUnauthorized
	}

	input := &dto.CreateTweetInput{
		Text: request.Text,
	}
	dto, err := h.tweetService.CreateTweet(user.Username(), input)

	if err != nil {
		return nil, err
	}

	return NewTweetResponse(dto), nil
}

func NewTweetResponse(tweet *dto.TweetDto) *pb.TweetResponse {
	return &pb.TweetResponse{
		Id:        tweet.ID,
		Text:      tweet.Text,
		CreatedAt: tweet.CreatedAt.String(),
		Author: &pb.AuthorResponse{
			ID:       tweet.Author.ID,
			Username: tweet.Author.Username,
		},
	}
}

func NewGrpcCreateTweetHandler(tweetService *services.TweetService) *GrpcCreateTweetHandler {
	return &GrpcCreateTweetHandler{
		tweetService: tweetService,
	}
}
