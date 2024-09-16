package main

import (
	"encoding/json"
	"hello-world-post/models"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/yuki5155/go-lambda-microkit/utils"
)

func TestHandler(t *testing.T) {
	testCases := []struct {
		name               string
		request            events.APIGatewayProxyRequest
		expectedBody       models.ResponseBody
		expectedStatusCode int
		expectedError      error
	}{
		{
			name: "valid request",
			request: events.APIGatewayProxyRequest{
				RequestContext: events.APIGatewayProxyRequestContext{
					Identity: events.APIGatewayRequestIdentity{
						SourceIP: "",
					},
				},
				Body: func() string {
					body := models.RequestBody{
						TestBody: "Hello, world!!",
					}
					bodyJson, _ := utils.BodyToJSON[models.RequestBody](body)
					return string(bodyJson)
				}(),
			},
			expectedBody: models.ResponseBody{
				Response: "Hello, world!",
			},
			expectedStatusCode: 200,
			expectedError:      nil,
		},
		{
			name: "localhost IP",
			request: events.APIGatewayProxyRequest{
				RequestContext: events.APIGatewayProxyRequestContext{
					Identity: events.APIGatewayRequestIdentity{
						SourceIP: "127.0.0.1",
					},
				},
				Body: "{}",
			},
			expectedBody: models.ResponseBody{
				Response: "Hello, world!",
			},
			expectedStatusCode: 400,
			expectedError:      nil,
		},
		{
			name: "invalid JSON",
			request: events.APIGatewayProxyRequest{
				Body: "{invalid json}",
			},
			expectedBody:       models.ResponseBody{},
			expectedStatusCode: 400,
			expectedError:      nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			response, err := handler(tc.request)

			if (err != nil) != (tc.expectedError != nil) {
				t.Errorf("[testcase:%s]handler() error = %v, expectedError %v", tc.name, err, tc.expectedError)
				return
			}

			if response.StatusCode != tc.expectedStatusCode {
				t.Errorf("[testcase:%s]handler() statusCode = %v, want %v", tc.name, response.StatusCode, tc.expectedStatusCode)
			}

			var gotBody models.ResponseBody
			if tc.expectedStatusCode == 200 {
				if err := json.Unmarshal([]byte(response.Body), &gotBody); err != nil {
					t.Errorf("[testcase:%s]failed to unmarshal response body: %v", tc.name, err)
					return
				}
			}
		})
	}
}
