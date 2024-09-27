package myaws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

type CognitoAPI interface {
	SignUp(ctx context.Context, params *cognitoidentityprovider.SignUpInput, optFns ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.SignUpOutput, error)
	ConfirmSignUp(ctx context.Context, params *cognitoidentityprovider.ConfirmSignUpInput, optFns ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ConfirmSignUpOutput, error)
}

type CognitoClientInterface interface {
	SignUp(ctx context.Context, username, password string) (string, error)
	ConfirmSignUp(ctx context.Context, username, confirmationCode string) error
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
	// Existing SignUp method implementation
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

func (c *cognitoClient) ConfirmSignUp(ctx context.Context, username, confirmationCode string) error {
	input := &cognitoidentityprovider.ConfirmSignUpInput{
		ClientId:         &c.ClientID,
		Username:         &username,
		ConfirmationCode: &confirmationCode,
	}

	_, err := c.API.ConfirmSignUp(ctx, input)
	return err
}
