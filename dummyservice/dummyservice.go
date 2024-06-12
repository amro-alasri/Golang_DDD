package dummyservice

import (
	"context"
	"learn/domain/user"
)

type Dummyservice struct {
}

// DTO (Data Transfer Object) which hold primative fields.
type FooReq struct {
	Username string
	Email    string
}

type SignupResp struct {
	ID string
}

func (d *Dummyservice) SignUp(ctx context.Context, username, email string) (*SignupResp, error) {
	userName, err := user.NewUsername(username)
	if err != nil {
		return nil, err
	}

	userEmail, err := user.NewEmail(email)
	if err != nil {
		return nil, err
	}

	// create the user
	new_user := user.New(userEmail, userName)

	// pass to repos

	return &SignupResp{
		ID: new_user.ID.String(),
	}, nil
}
