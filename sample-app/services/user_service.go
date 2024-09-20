package services

import "github.com/yuki5155/go-lambda-microkit/domains"

// /root/go/bin/mockgen -source=services/user_service.go -destination=mocks/mock_user_service.go -package=mocks
type IUserService interface {
	GetUser() (domains.User, error)
	GetEmail() (string, error)
}

type UserService struct {
}

func (u *UserService) GetUser() (domains.User, error) {
	return domains.User{
		Name:  "User",
		Email: "Email",
	}, nil
}

func (u *UserService) GetEmail() (string, error) {
	return "Email", nil
}

func NewUserService() IUserService {
	return &UserService{}
}

var _ IUserService = (*UserService)(nil)
