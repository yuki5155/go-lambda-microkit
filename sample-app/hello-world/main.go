package main

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/yuki5155/go-lambda-microkit/myaws/lambda/middleware"
	"github.com/yuki5155/go-lambda-microkit/services"
)

var NewUserService func() services.IUserService = services.NewUserService

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Received request: %+v", request)

	userService := NewUserService()
	user, err := userService.GetUser()
	if err != nil {
		log.Printf("Error getting user: %v", err)
		return events.APIGatewayProxyResponse{
			Body:       err.Error(),
			StatusCode: 500,
		}, nil
	}

	response := map[string]interface{}{
		"token": "mock-token-12345",
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
