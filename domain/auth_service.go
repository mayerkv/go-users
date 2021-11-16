package domain

import "context"

type AuthService interface {
	CreateAccount(ctx context.Context, email, password, userId string, role UserRole) error
}
