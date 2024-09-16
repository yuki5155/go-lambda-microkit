package main

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/yuki5155/go-lambda-microkit/utils"
)

type Body struct {
	TestBody string `json:"testBody"` // JSONのキーに合わせて小文字に変更
}

type ResponseBody struct {
	Response string `json:"response"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	body, err := utils.NewRequestBody[Body](request, Body{})
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Bad Request",
		}, nil
	}

	if body.TestBody == "" {
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Bad Request",
		}, nil
	}

	resBody := ResponseBody{
		Response: fmt.Sprintf("Received TestBody: %s", body.TestBody),
	}

	resBodyJson, err := json.Marshal(resBody)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Internal Server Error",
		}, nil
	}

	return events.APIGatewayProxyResponse{
		Body:       string(resBodyJson),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
