package user

import (
	"errors"

	"github.com/google/uuid"
)

// list of errors
var (
	ErrInvalidID = errors.New("invalid id provided")
)

// ID
type ID string

func NewID(i string) (ID, error) {
	_, err := uuid.Parse(i)
	if err != nil {
		return "", errors.Join(ErrInvalidID, err)
	}

	return ID(i), nil
}

func (i ID) String() string {
	return string(i)
}

// Email
type Email string

func NewEmail(i string) (Email, error) {
	// domain logic comes in
	//vallidate it is an valid email
	return Email(i), nil
}

func (i Email) String() string {

	return string(i)
}

// Username
type Username string

func NewUsername(i string) (Username, error) {
	// domain logic comes in
	//vallidate it is an valid email
	return Username(i), nil
}

func (i Username) String() string {
	return string(i)
}

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
