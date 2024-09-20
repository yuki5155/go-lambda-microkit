package main

import (
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/golang/mock/gomock"
	"github.com/yuki5155/go-lambda-microkit/mocks"
	"github.com/yuki5155/go-lambda-microkit/services"
)

func TestHandler(t *testing.T) {
	testCases := []struct {
		name          string
		request       events.APIGatewayProxyRequest
		expectedBody  string
		expectedError error
		mockUser      string
		mockEmail     string
		setupMock     func(*mocks.MockIUserService)
	}{
		{
			name: "empty IP",
			request: events.APIGatewayProxyRequest{
				RequestContext: events.APIGatewayProxyRequestContext{
					Identity: events.APIGatewayRequestIdentity{
						SourceIP: "",
					},
				},
			},
			expectedBody: "Hello, world!\n",
			mockUser:     "MockUser1",
			mockEmail:    "mock1@example.com",
			setupMock: func(mockUserService *mocks.MockIUserService) {
				mockUserService.EXPECT().GetUser().Return("MockUser1").Times(1)
				mockUserService.EXPECT().GetEmail().Return("mock1@example.com").Times(1)
			},
		},
		{
			name: "localhost IP",
			request: events.APIGatewayProxyRequest{
				RequestContext: events.APIGatewayProxyRequestContext{
					Identity: events.APIGatewayRequestIdentity{
						SourceIP: "127.0.0.1",
					},
				},
			},
			expectedBody: "Hello, 127.0.0.1!\n",
			mockUser:     "MockUser2",
			mockEmail:    "mock2@example.com",
			setupMock: func(mockUserService *mocks.MockIUserService) {
				mockUserService.EXPECT().GetUser().Return("MockUser2").Times(1)
				mockUserService.EXPECT().GetEmail().Return("mock2@example.com").Times(1)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUserService := mocks.NewMockIUserService(ctrl)
			testCase.setupMock(mockUserService)

			origNewUserService := NewUserService
			NewUserService = func() services.IUserService {
				return mockUserService
			}
			defer func() { NewUserService = origNewUserService }()

			response, err := handler(testCase.request)
			if err != testCase.expectedError {
				t.Errorf("Expected error %v, but got %v", testCase.expectedError, err)
			}

			if response.Body != testCase.expectedBody {
				t.Errorf("Expected response %v, but got %v", testCase.expectedBody, response.Body)
			}

			if response.StatusCode != 200 {
				t.Errorf("Expected status code 200, but got %v", response.StatusCode)
			}
		})
	}
}
