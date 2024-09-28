package myaws_test

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/joho/godotenv"
	"github.com/yuki5155/go-lambda-microkit/myaws"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

func setupCognitoClient(t *testing.T) myaws.CognitoClientInterface {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		t.Fatalf("failed to load configuration, %v", err)
	}
	client := cognitoidentityprovider.NewFromConfig(cfg)
	if client == nil {
		t.Fatalf("failed to create Cognito client")
	}
	cloudformationConf := cloudformation.NewFromConfig(cfg)

	cloudFormationClient := myaws.NewCloudFormationClient(cloudformationConf)
	userPoolID, err := cloudFormationClient.GetCloudFormationOutput(context.Background(), "cognito", "UserPoolId")
	if err != nil {
		t.Fatalf("Failed to get CloudFormation output: %v", err)
	}
	userPoolClientId, err := cloudFormationClient.GetCloudFormationOutput(context.Background(), "cognito", "UserPoolClientId")
	if err != nil {
		t.Fatalf("Failed to get CloudFormation output: %v", err)
	}

	return myaws.NewCognitoClient(client, userPoolClientId, userPoolID)
}

func TestCognitoSignUp(t *testing.T) {
	cognitoClient := setupCognitoClient(t)
	username := os.Getenv("SAMPLEEMAILADDRESS")
	password := os.Getenv("SAMPLEPASSWORD")
	userSub, err := cognitoClient.SignUp(context.Background(), username, password)
	if err != nil {
		t.Errorf("Failed to sign up: %v", err)
	} else {
		t.Logf("UserSub: %s", userSub)
	}
}

func TestCognitoConfirmSignUp(t *testing.T) {
	cognitoClient := setupCognitoClient(t)
	username := os.Getenv("SAMPLEEMAILADDRESS")
	confirmationCode := ""

	err := cognitoClient.ConfirmSignUp(context.Background(), username, confirmationCode)
	if err != nil {
		t.Errorf("Failed to confirm sign up: %v", err)
	} else {
		t.Logf("Successfully confirmed sign up for user: %s", username)
	}
}

func TestCognitoLogin(t *testing.T) {
	cognitoClient := setupCognitoClient(t)
	username := os.Getenv("SAMPLEEMAILADDRESS")
	password := os.Getenv("SAMPLEPASSWORD")

	accessToken, err := cognitoClient.Login(context.Background(), username, password)
	if err != nil {
		t.Errorf("Failed to login: %v", err)
	} else {
		t.Logf("Successfully logged in. Access Token: %s", accessToken)
	}
}

func TestCognitoLogout(t *testing.T) {
	cognitoClient := setupCognitoClient(t)
	username := os.Getenv("SAMPLEEMAILADDRESS")
	password := os.Getenv("SAMPLEPASSWORD")

	// First, login to get an access token
	accessToken, err := cognitoClient.Login(context.Background(), username, password)
	if err != nil {
		t.Fatalf("Failed to login: %v", err)
	}

	// Now, attempt to logout
	err = cognitoClient.Logout(context.Background(), accessToken)
	if err != nil {
		t.Errorf("Failed to logout: %v", err)
	} else {
		t.Logf("Successfully logged out")
	}

}
