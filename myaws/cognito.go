package myaws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

type CognitoAPI interface {
	SignUp(ctx context.Context, params *cognitoidentityprovider.SignUpInput, optFns ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.SignUpOutput, error)
}

type CognitoClientInterface interface {
	SignUp(ctx context.Context, username, password string) (string, error)
}

type cognitoClient struct {
	API        CognitoAPI
	ClientID   string
	UserPoolID string
}

func NewCognitoClient(api CognitoAPI, clientID, userPoolID string) CognitoClientInterface {
	return &cognitoClient{
		API:        api,
		ClientID:   clientID,
		UserPoolID: userPoolID,
	}
}

func (c *cognitoClient) SignUp(ctx context.Context, username, password string) (string, error) {
	// username should be an email address
	input := &cognitoidentityprovider.SignUpInput{
		ClientId: &c.ClientID,
		Username: &username,
		Password: &password,
	}

	result, err := c.API.SignUp(ctx, input)
	if err != nil {
		return "", err
	}

	return *result.UserSub, nil
}
