package domain

import "errors"

var (
	ErrUserAlreadyExists = errors.New("user already exists")
)

type UserService struct {
	userRepository UserRepository
	authService    AuthService
}

func NewUserService(userRepository UserRepository, authService AuthService) *UserService {
	return &UserService{userRepository: userRepository, authService: authService}
}

func (s *UserService) CreateUser(email, password string, role UserRole) (*User, error) {
	u, err := s.userRepository.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	if u != nil {
		return nil, ErrUserAlreadyExists
	}

	user := CreateUser(email, role)

	if err := s.userRepository.Save(user); err != nil {
		return nil, err
	}

	if err := s.authService.CreateAccount(email, password, user.Id, user.Role); err != nil {
		s.userRepository.Delete(user)
		return nil, err
	}

	return user, nil
}
