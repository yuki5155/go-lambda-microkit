package main

import (
	"fmt"

	"hello-world-post/models"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/yuki5155/go-lambda-microkit/utils"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	body, err := utils.NewRequestBody[models.RequestBody](request, models.RequestBody{})
	if err != nil {
		return utils.BadRequestResponse()
	}

	if body.TestBody == "" {
		return utils.BadRequestResponse("TestBody cannot be empty")
	}

	resBody := models.ResponseBody{
		Response: fmt.Sprintf("Received TestBody: %s", body.TestBody),
	}

	return utils.SuccessResponse(resBody)
}

func main() {
	lambda.Start(handler)
}
