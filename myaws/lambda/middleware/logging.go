package middleware

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
)

// LoggingMiddleware returns a Middleware that logs requests and responses
func LoggingMiddleware() Middleware {
	return func(next LambdaHandler) LambdaHandler {
		return func(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
			// Log request
			reqJSON, _ := json.Marshal(req)
			log.Printf("Request: %s", string(reqJSON))

			// Call the next handler
			resp, err := next(req)

			// Log response
			respJSON, _ := json.Marshal(resp)
			log.Printf("Response: %s", string(respJSON))

			return resp, err
		}
	}
}
