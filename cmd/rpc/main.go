package main

import (
	"context"
	"log"
	"net"

	grpcHandler "github.com/harlancleiton/go-tweets/handlers/grpc"
	"github.com/harlancleiton/go-tweets/internal/application/services"
	"github.com/harlancleiton/go-tweets/internal/infra/persistence/memory"
	grpcInterceptor "github.com/harlancleiton/go-tweets/middlewares/grpc"
	"github.com/harlancleiton/go-tweets/pkg/pb"
	"google.golang.org/grpc"
)

func main() {
	authInterceptor := grpcInterceptor.NewAuthInterceptor(memory.NewMemoryUserRepository())
	options := []grpc.ServerOption{grpc.UnaryInterceptor(authInterceptor.Unary())}
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

type TweetService struct {
	pb.UnimplementedTweetServiceServer
	createHandler *grpcHandler.GrpcCreateTweetHandler
}

func (s *TweetService) Create(ctx context.Context, request *pb.CreateTweetRequest) (*pb.TweetResponse, error) {
	return s.createHandler.Create(ctx, request)
}

func registerTweetServiceServer(server *grpc.Server) {
	service := services.NewTweetService(memory.NewMemoryUserRepository(), memory.NewMemoryTweetRepository())
	pb.RegisterTweetServiceServer(server, &TweetService{
		createHandler: grpcHandler.NewGrpcCreateTweetHandler(service),
	})
}
