package services

import (
	"context"
	"github.com/mayerkv/go-auth/grpc-service"
	"github.com/mayerkv/go-users/domain"
	"time"
)

type AuthService struct {
	authServiceClient grpc_service.AuthServiceClient
}

func NewAuthService(authServiceClient grpc_service.AuthServiceClient) *AuthService {
	return &AuthService{authServiceClient: authServiceClient}
}

func (s *AuthService) CreateAccount(email, password, userId string, role domain.UserRole) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var accountRole grpc_service.AccountRole
	switch role {
	case domain.UserRoleUser:
		accountRole = grpc_service.AccountRole_USER
	case domain.UserRoleAdmin:
		accountRole = grpc_service.AccountRole_ADMIN
	}

	req := &grpc_service.CreateAccountRequest{
		Email:    email,
		Password: password,
		UserId:   userId,
		Role:     accountRole,
	}

	_, err := s.authServiceClient.CreateAccount(ctx, req)

	return err
}
