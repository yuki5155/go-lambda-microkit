package utils

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

type DefaultBody struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

const (
	StatusBadRequestDefaultMessage          = "The request could not be understood by the server due to malformed syntax."
	StatusUnauthorizedDefaultMessage        = "The request requires user authentication."
	StatusNotFoundDefaultMessage            = "The requested resource could not be found."
	StatusInternalServerErrorDefaultMessage = "The server encountered an unexpected condition which prevented it from fulfilling the request."
)

const (
	StatusBadRequestDefaultCode          = "BAD_REQUEST"
	StatusUnauthorizedDefaultCode        = "UNAUTHORIZED"
	StatusNotFoundDefaultCode            = "NOT_FOUND"
	StatusInternalServerErrorDefaultCode = "INTERNAL_SERVER_ERROR"
)

var defaultErrorMessages = map[int]DefaultBody{
	http.StatusBadRequest: {
		Message: StatusBadRequestDefaultMessage,
		Code:    StatusBadRequestDefaultCode,
	},
	http.StatusUnauthorized: {
		Message: StatusUnauthorizedDefaultMessage,
		Code:    StatusUnauthorizedDefaultCode,
	},
	http.StatusNotFound: {
		Message: StatusNotFoundDefaultMessage,
		Code:    StatusNotFoundDefaultCode,
	},
	http.StatusInternalServerError: {
		Message: StatusInternalServerErrorDefaultMessage,
		Code:    StatusInternalServerErrorDefaultCode,
	},
}

// SuccessResponse creates a successful response (status code 200) with a required body
func SuccessResponse(body interface{}) (events.APIGatewayProxyResponse, error) {
	return createJSONResponse(http.StatusOK, body)
}

// BadRequestResponse creates a bad request response (status code 400) with an optional message
func BadRequestResponse(message ...string) (events.APIGatewayProxyResponse, error) {
	return createErrorResponse(http.StatusBadRequest, message...)
}

// InternalServerErrorResponse creates an internal server error response (status code 500) with an optional message
func InternalServerErrorResponse(message ...string) (events.APIGatewayProxyResponse, error) {
	return createErrorResponse(http.StatusInternalServerError, message...)
}

// NotFoundResponse creates a not found response (status code 404) with an optional message
func NotFoundResponse(message ...string) (events.APIGatewayProxyResponse, error) {
	return createErrorResponse(http.StatusNotFound, message...)
}

// UnauthorizedResponse creates an unauthorized response (status code 401) with an optional message
func UnauthorizedResponse(message ...string) (events.APIGatewayProxyResponse, error) {
	return createErrorResponse(http.StatusUnauthorized, message...)
}

// createErrorResponse is a helper function to create an error response
func createErrorResponse(statusCode int, message ...string) (events.APIGatewayProxyResponse, error) {
	body := defaultErrorMessages[statusCode]
	if len(message) > 0 && message[0] != "" {
		body.Message = message[0]
	}
	return createJSONResponse(statusCode, body)
}

// createJSONResponse is a helper function to create a JSON response
func createJSONResponse(statusCode int, body interface{}) (events.APIGatewayProxyResponse, error) {
	bodyJSON, err := BodyToJSON(body)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       bodyJSON,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}
