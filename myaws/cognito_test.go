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
	// .envファイルを読み込む
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

func TestCognitoSIgnUp(t *testing.T) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		t.Errorf("failed to load configuration, %v", err)
	}
	client := cognitoidentityprovider.NewFromConfig(cfg)
	if client == nil {
		t.Errorf("failed to create Cognito client")
	}
	cloudformationConf := cloudformation.NewFromConfig(cfg)

	cloudFormationClient := myaws.NewCloudFormationClient(cloudformationConf)
	userPoolID, err := cloudFormationClient.GetCloudFormationOutput(context.Background(), "cognito", "UserPoolId")
	if err != nil {
		t.Errorf("Failed to get CloudFormation output: %v", err)
	} else {
		t.Logf("UserPoolId: %s", userPoolID)
	}
	userPoolClientId, err := cloudFormationClient.GetCloudFormationOutput(context.Background(), "cognito", "UserPoolClientId")
	if err != nil {
		t.Errorf("Failed to get CloudFormation output: %v", err)
	} else {
		t.Logf("UserPoolClientId: %s", userPoolClientId)
	}

	cognitoClient := myaws.NewCognitoClient(client, userPoolClientId, userPoolID)
	username := os.Getenv("SAMPLEEMAILADDRESS")
	password := os.Getenv("SAMPLEPASSWORD")
	userSub, err := cognitoClient.SignUp(context.Background(), username, password)
	if err != nil {
		t.Errorf("Failed to sign up: %v", err)
	} else {
		t.Logf("UserSub: %s", userSub)
	}

}
