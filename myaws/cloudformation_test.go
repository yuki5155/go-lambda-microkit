package myaws_test

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/joho/godotenv"
	"github.com/yuki5155/go-lambda-microkit/myaws"
)

func init() {
	// .envファイルを読み込む
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

func TestOutput(t *testing.T) {

	// build the client of CloudFormation
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		t.Errorf("failed to load configuration, %v", err)
	}

	client := cloudformation.NewFromConfig(cfg)
	if client == nil {
		t.Errorf("failed to create CloudFormation client")
	}

	cloudFormationClient := myaws.NewCloudFormationClient(client)
	poolid, err := cloudFormationClient.GetCloudFormationOutput(context.Background(), "cognito", "UserPoolId")
	if err != nil {
		t.Errorf("Failed to get CloudFormation output: %v", err)
	} else {
		t.Logf("UserPoolId: %s", poolid)
	}
	fmt.Println(poolid)
}
