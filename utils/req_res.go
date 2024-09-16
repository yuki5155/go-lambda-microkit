package utils

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

func NewRequestBody[T any](proxyRequest events.APIGatewayProxyRequest, additionalData T) (T, error) {
	if proxyRequest.Body == "" {
		return additionalData, nil
	}
	if err := json.Unmarshal([]byte(proxyRequest.Body), &additionalData); err != nil {
		return additionalData, err
	}
	return additionalData, nil
}
