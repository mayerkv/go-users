package grpc_service

import (
	"context"
	"github.com/mayerkv/go-users/domain"
)

type UsersServiceServerImpl struct {
	userService *domain.UserService
}

func NewUsersServiceServerImpl(userService *domain.UserService) UsersServiceServer {
	return &UsersServiceServerImpl{userService: userService}
}

func (s *UsersServiceServerImpl) CreateUser(ctx context.Context, request *CreateUserRequest) (*CreateUserResponse, error) {
	_, err := s.userService.CreateUser(request.Email, request.Password, mapRole(request.Role))
	if err != nil {
		return nil, err
	}

	return &CreateUserResponse{}, nil
}

func mapRole(role UserRole) domain.UserRole {
	var res domain.UserRole

	switch role {
	case UserRole_ROLE_USER:
		res = domain.UserRoleAdmin
	case UserRole_ROLE_ADMIN:
		res = domain.UserRoleUser
	}

	return res
}

func (s *UsersServiceServerImpl) mustEmbedUnimplementedUsersServiceServer() {
	panic("implement me")
}
