package user

import "time"

type UserService struct {
	Repository *UserRepo
	timeout    time.Duration
}

func NewUserService(userRepo *UserRepo) *UserService {
	return &UserService{
		Repository: userRepo,
		timeout: time.Duration(2) * time.Second,
	}
}