package domain

import "github.com/google/uuid"

type UserRole int

const (
	UserRoleUser UserRole = iota
	UserRoleAdmin
)

type User struct {
	Id    string
	Email string
	Role  UserRole
}

func CreateUser(email string, role UserRole) *User {
	return &User{
		Id:    uuid.NewString(),
		Email: email,
		Role:  role,
	}
}
