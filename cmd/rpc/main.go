package main

import (
	"context"
	"log"
	"net"

	"github.com/harlancleiton/go-tweets/internal/application/services"
	"github.com/harlancleiton/go-tweets/internal/domain/dto"
	infra "github.com/harlancleiton/go-tweets/internal/infra/persistence/memory"
	"github.com/harlancleiton/go-tweets/pkg/pb"
	"google.golang.org/grpc"
)

type TweetServiceServerAdapter struct {
	tweetService *services.TweetService
	pb.UnimplementedTweetServiceServer
}

func (adapter *TweetServiceServerAdapter) Create(ctx context.Context, request *pb.CreateTweetRequest) (*pb.TweetResponse, error) {
	input := &dto.CreateTweetInput{
		Text: request.Text,
	}
	tweet, err := adapter.tweetService.CreateTweet("username", input)

	if err != nil {
		return nil, err
	}

	return &pb.TweetResponse{
		Id:        tweet.ID,
		Text:      tweet.Text,
		CreatedAt: tweet.CreatedAt.String(),
	}, nil
}

func NewTweetServiceServerAdapter(tweetService *services.TweetService) *TweetServiceServerAdapter {
	return &TweetServiceServerAdapter{
		tweetService: tweetService,
	}
}

func registerTweetServiceServer(server *grpc.Server) {
	service := services.NewTweetService(infra.NewMemoryUserRepository(), infra.NewMemoryTweetRepository())
	pb.RegisterTweetServiceServer(server, NewTweetServiceServerAdapter(service))
}

func main() {
	options := []grpc.ServerOption{}
	server := grpc.NewServer(options...)

	registerTweetServiceServer(server)

	listener, err := net.Listen("tcp", "localhost:50051")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
