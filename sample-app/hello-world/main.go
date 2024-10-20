package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/yuki5155/go-lambda-microkit/myaws"
	"github.com/yuki5155/go-lambda-microkit/myaws/lambda/middleware"
	"github.com/yuki5155/go-lambda-microkit/services"
)

var NewUserService func(cognitoClient myaws.CognitoClientInterface) services.IUserService = services.NewUserService

var cognitoClient myaws.CognitoClientInterface

func init() {
	// Load the AWS SDK configuration
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		log.Fatalf("Failed to load AWS config: %v", err)
	}

	// Create the Cognito service client
	cognitoAPI := cognitoidentityprovider.NewFromConfig(cfg)

	// Get the Client ID and User Pool ID from environment variables
	clientID := os.Getenv("COGNITO_CLIENT_ID")
	userPoolID := os.Getenv("COGNITO_USER_POOL_ID")

	if clientID == "" || userPoolID == "" {
		log.Println("COGNITO_CLIENT_ID and/or COGNITO_USER_POOL_ID are not set")
		log.Printf("COGNITO_CLIENT_ID set: %v", clientID != "")
		log.Printf("COGNITO_USER_POOL_ID set: %v", userPoolID != "")
		// Instead of fatal, we'll continue with nil cognitoClient
		cognitoClient = nil
	} else {
		// Initialize the Cognito client
		cognitoClient = myaws.NewCognitoClient(cognitoAPI, clientID, userPoolID)
		log.Println("Cognito client initialized successfully")
	}
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Received request: %+v", request)

	if cognitoClient == nil {
		log.Println("Cognito client is not initialized")
		return events.APIGatewayProxyResponse{
			Body:       "Internal Server Error: Cognito client not initialized",
			StatusCode: 500,
		}, nil
	}

	userService := NewUserService(cognitoClient)
	user, err := userService.GetUser()
	if err != nil {
		log.Printf("Error getting user: %v", err)
		return events.APIGatewayProxyResponse{
			Body:       fmt.Sprintf("Error getting user: %v", err),
			StatusCode: 500,
		}, nil
	}

	response := map[string]interface{}{
		"token": "mock-token-123456",
		"user": map[string]string{
			"name":  user.Name,
			"email": user.Email,
		},
	}

	responseBody, err := json.Marshal(response)
	if err != nil {
		log.Printf("Error marshaling response: %v", err)
		return events.APIGatewayProxyResponse{
			Body:       "Internal Server Error",
			StatusCode: 500,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		Body:       string(responseBody),
		StatusCode: 200,
	}, nil
}

func main() {
	handlerWithMiddleware := middleware.Chain(
		handler,
		middleware.CORSMiddleware(),
		middleware.LoggingMiddleware(),
	)
	lambda.Start(handlerWithMiddleware)
}
