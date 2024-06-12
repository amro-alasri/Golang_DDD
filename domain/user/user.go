package user

import "github.com/google/uuid"

type ID string
type Email string
type Username string

type User struct {
	ID       ID
	Username Username
	Email    Email
}

func New(email Email, username Username) User {
	return User{
		ID:       ID(uuid.NewString()),
		Email:    email,
		Username: username,
	}
}
