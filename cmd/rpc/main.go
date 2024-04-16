package main

import (
	"context"
	"log"
	"net"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	grpcHandler "github.com/harlancleiton/go-tweets/handlers/grpc"
	"github.com/harlancleiton/go-tweets/internal/application/services"
	"github.com/harlancleiton/go-tweets/internal/infra/persistence/memory"
	"github.com/harlancleiton/go-tweets/pkg/domain/events"
	grpcInterceptor "github.com/harlancleiton/go-tweets/pkg/middlewares/grpc"
	"github.com/harlancleiton/go-tweets/pkg/pb"
	"google.golang.org/grpc"
)

func main() {
	generateAccessToken()

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
	service := services.NewTweetService(memory.NewMemoryUserRepository(), memory.NewMemoryTweetRepository(), events.NewConcreteEventDispatcher())
	pb.RegisterTweetServiceServer(server, &TweetService{
		createHandler: grpcHandler.NewGrpcCreateTweetHandler(service),
	})
}

func generateAccessToken() {
	claims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Minute * 5).Unix(),
		Issuer:    "tweets",
		Subject:   "somebody",
		Audience:  strings.Join(([]string{"somebody_else"}), ", "),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	mySigningKey := []byte("secret")
	ss, err := token.SignedString(mySigningKey)

	if err != nil {
		log.Fatalf("Failed to generate token: %v", err)
	}

	log.Println("generated token:", ss)
}
