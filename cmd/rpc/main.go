package main

import (
	"context"
	"log"
	"net"

	grpcHandler "github.com/harlancleiton/go-tweets/handlers/grpc"
	"github.com/harlancleiton/go-tweets/internal/application/services"
	"github.com/harlancleiton/go-tweets/internal/infra/persistence/memory"
	grpcInterceptor "github.com/harlancleiton/go-tweets/pkg/middlewares/grpc"
	"github.com/harlancleiton/go-tweets/pkg/pb"
	"google.golang.org/grpc"
)

func main() {
	log.Println("Starting gRPC server...")

	authInterceptor := grpcInterceptor.NewAuthInterceptor(memory.NewMemoryUserRepository())
	options := []grpc.ServerOption{grpc.UnaryInterceptor(authInterceptor.Unary())}
	server := grpc.NewServer(options...)

	log.Println("gRPC server started")

	registerTweetServiceServer(server)

	log.Println("TweetService registered")
	log.Println("All services registered")

	address := "0.0.0.0:50051"
	listener, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Println("Listening on", address)

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
