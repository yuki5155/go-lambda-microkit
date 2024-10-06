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
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		log.Fatalf("Failed to load AWS config: %v", err)
	}

	cognitoAPI := cognitoidentityprovider.NewFromConfig(cfg)

	fmt.Println("COGNITO_CLIENT_ID: ", os.Getenv("COGNITO_CLIENT_ID"))

	clientID := os.Getenv("COGNITO_CLIENT_ID")
	userPoolID := os.Getenv("COGNITO_USER_POOL_ID")

	if clientID == "" || userPoolID == "" {
		log.Fatalf("COGNITO_CLIENT_ID and/or COGNITO_USER_POOL_ID are not set")
	}

	cognitoClient = myaws.NewCognitoClient(cognitoAPI, clientID, userPoolID)
	log.Println("Cognito client initialized successfully")
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Received login request: %+v", request)

	var loginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := json.Unmarshal([]byte(request.Body), &loginRequest)
	if err != nil {
		log.Printf("Error unmarshaling request body: %v", err)
		return events.APIGatewayProxyResponse{
			Body:       "Invalid request body",
			StatusCode: 400,
		}, nil
	}

	userService := NewUserService(cognitoClient)
	user, err := userService.Login(context.Background(), loginRequest.Username, loginRequest.Password)
	if err != nil {
		log.Printf("Login error: %v", err)
		return events.APIGatewayProxyResponse{
			Body:       "Login failed",
			StatusCode: 401,
		}, nil
	}

	response := map[string]interface{}{
		"token": user.IDToken,
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
