package user

import (
	"context"
	"time"
)

type service struct {
	Repository *IRepository
	timeout    time.Duration
}

func NewUserService(userRepo *IRepository) IService {
	return &service{
		Repository: userRepo,
		timeout:    time.Duration(2) * time.Second,
	}
}

func (service *service) SignIn(ctx context.Context, req *CreateUserReq) (*CreateUserRes, error) {
	//TODO implement me
	panic("implement me")
}

func (service *service) SingUp(ctx context.Context, req *LoginUserReq) (*LoginUserRes, error) {
	//TODO implement me
	panic("implement me")
}
