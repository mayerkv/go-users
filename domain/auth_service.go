package domain

type AuthService interface {
	CreateAccount(email, password, userId string, role UserRole) error
}
