package services

import (
	"context"

	"github.com/yuki5155/go-lambda-microkit/domains"
	"github.com/yuki5155/go-lambda-microkit/myaws"
)

// /root/go/bin/mockgen -source=services/user_service.go -destination=mocks/mock_user_service.go -package=mocks
type IUserService interface {
	GetUser() (domains.User, error)
	GetEmail() (string, error)
	Login(ctx context.Context, username, password string) (domains.User, error)
}

type UserService struct {
	cognitoClient myaws.CognitoClientInterface
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
func (u *UserService) Login(ctx context.Context, username, password string) (domains.User, error) {
	_, idToken, err := u.cognitoClient.Login(ctx, username, password)
	if err != nil {
		return domains.User{}, err
	}

	user := domains.User{
		Name:    username,
		Email:   username,
		IDToken: idToken,
	}

	return user, nil
}

func NewUserService(cognitoClient myaws.CognitoClientInterface) IUserService {
	return &UserService{
		cognitoClient: cognitoClient,
	}
}

var _ IUserService = (*UserService)(nil)
