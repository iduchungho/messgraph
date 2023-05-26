package user

import (
	"context"
)

type IRepository interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
}

type IService interface {
	SignIn(ctx context.Context, req *CreateUserReq) (*CreateUserRes, error)
	SingUp(ctx context.Context, req *LoginUserReq) (*LoginUserRes, error)
}
