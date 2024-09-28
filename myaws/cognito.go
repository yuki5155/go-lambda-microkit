package myaws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

type CognitoAPI interface {
	SignUp(ctx context.Context, params *cognitoidentityprovider.SignUpInput, optFns ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.SignUpOutput, error)
	ConfirmSignUp(ctx context.Context, params *cognitoidentityprovider.ConfirmSignUpInput, optFns ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ConfirmSignUpOutput, error)
	InitiateAuth(ctx context.Context, params *cognitoidentityprovider.InitiateAuthInput, optFns ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.InitiateAuthOutput, error)
	GlobalSignOut(ctx context.Context, params *cognitoidentityprovider.GlobalSignOutInput, optFns ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.GlobalSignOutOutput, error)
}

type CognitoClientInterface interface {
	SignUp(ctx context.Context, username, password string) (string, error)
	ConfirmSignUp(ctx context.Context, username, confirmationCode string) error
	Login(ctx context.Context, username, password string) (string, string, error) // Updated to return two strings
	Logout(ctx context.Context, accessToken string) error
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

func (c *cognitoClient) Login(ctx context.Context, username, password string) (string, string, error) {
	input := &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow: "USER_PASSWORD_AUTH",
		AuthParameters: map[string]string{
			"USERNAME": username,
			"PASSWORD": password,
		},
		ClientId: aws.String(c.ClientID),
	}

	result, err := c.API.InitiateAuth(ctx, input)
	if err != nil {
		return "", "", err
	}

	if result.AuthenticationResult == nil {
		return "", "", fmt.Errorf("no authentication result")
	}

	return *result.AuthenticationResult.AccessToken, *result.AuthenticationResult.IdToken, nil
}

func (c *cognitoClient) Logout(ctx context.Context, accessToken string) error {
	input := &cognitoidentityprovider.GlobalSignOutInput{
		AccessToken: aws.String(accessToken),
	}

	_, err := c.API.GlobalSignOut(ctx, input)
	return err
}
