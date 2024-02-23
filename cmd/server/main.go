package main

import (
	"context"
	"log"
	"net"

	"github.com/harlancleiton/go-tweets/internal/dto"
	"github.com/harlancleiton/go-tweets/internal/usecases"
	"github.com/harlancleiton/go-tweets/pkg/pb"
	"google.golang.org/grpc"
)

type TweetServiceServerAdapter struct {
	createTweet *usecases.CreateTweet
	pb.UnimplementedTweetServiceServer
}

func (adapter *TweetServiceServerAdapter) Create(ctx context.Context, request *pb.CreateTweetRequest) (*pb.TweetResponse, error) {
	input := dto.CreateTweetInput{
		Text: request.Text,
	}
	tweet, err := adapter.createTweet.Execute(ctx, input)

	if err != nil {
		return nil, err
	}

	return &pb.TweetResponse{
		Id:        tweet.ID,
		Text:      tweet.Text,
		CreatedAt: tweet.CreatedAt.String(),
	}, nil
}

func NewTweetServiceServerAdapter(createTweet *usecases.CreateTweet) *TweetServiceServerAdapter {
	return &TweetServiceServerAdapter{
		createTweet: createTweet,
	}
}

func main() {
	options := []grpc.ServerOption{}
	server := grpc.NewServer(options...)

	pb.RegisterTweetServiceServer(server, NewTweetServiceServerAdapter(usecases.NewCreateTweet()))

	listener, err := net.Listen("tcp", "localhost:50051")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
