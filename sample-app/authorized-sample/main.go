package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type CognitoClaims struct {
	Sub             string `json:"sub"`
	Email           string `json:"email"`
	CognitoUsername string `json:"cognito:username"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	claims, ok := request.RequestContext.Authorizer["claims"].(map[string]interface{})
	if !ok {
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Unable to parse Cognito claims",
		}, nil
	}
	fmt.Println(claims)
	return events.APIGatewayProxyResponse{
		Body:       "Hello, world!\n",
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
