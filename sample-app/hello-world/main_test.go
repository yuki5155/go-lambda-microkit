package main

import (
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/golang/mock/gomock"
	"github.com/yuki5155/go-lambda-microkit/domains"
	"github.com/yuki5155/go-lambda-microkit/mocks"
	"github.com/yuki5155/go-lambda-microkit/services"
)

type mockServices struct {
	userService *mocks.MockIUserService
	// Add other service mocks here as needed
}

func TestHandler(t *testing.T) {
	userDomain := domains.User{
		Name:  "User",
		Email: "Email",
	}
	testCases := []struct {
		name          string
		request       events.APIGatewayProxyRequest
		expectedBody  string
		expectedError error
		setupMocks    func(*mockServices)
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
			setupMocks: func(m *mockServices) {
				m.userService.EXPECT().GetUser().Return(userDomain, nil).Times(1)
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
			setupMocks: func(m *mockServices) {
				m.userService.EXPECT().GetUser().Return(userDomain, nil).Times(1)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mocks := &mockServices{
				userService: mocks.NewMockIUserService(ctrl),
				// Initialize other service mocks here
			}

			testCase.setupMocks(mocks)

			// Replace the original NewUserService function
			origNewUserService := NewUserService
			NewUserService = func() services.IUserService {
				return mocks.userService
			}
			defer func() { NewUserService = origNewUserService }()

			// Add similar replacements for other services here

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
