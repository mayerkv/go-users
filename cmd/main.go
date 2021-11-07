package main

import (
	auth_service "github.com/mayerkv/go-auth/grpc-service"
	"github.com/mayerkv/go-users/domain"
	"github.com/mayerkv/go-users/grpc-service"
	"github.com/mayerkv/go-users/repository"
	"github.com/mayerkv/go-users/services"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func main() {
	authAddr := os.Getenv("AUTH_ADDR")
	if authAddr == "" {
		authAddr = "auth:9090"
	}

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(authAddr, opts...)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	authServiceClient := auth_service.NewAuthServiceClient(conn)
	userRepository := repository.NewInMemoryUserRepository()
	authService := services.NewAuthService(authServiceClient)
	userService := domain.NewUserService(userRepository, authService)
	srv := grpc_service.NewUsersServiceServerImpl(userService)

	if err := runGrpcServer(srv); err != nil {
		log.Fatal(err)
	}
}

func runGrpcServer(srv grpc_service.UsersServiceServer) error {
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		return err
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	grpc_service.RegisterUsersServiceServer(grpcServer, srv)

	return grpcServer.Serve(lis)
}
