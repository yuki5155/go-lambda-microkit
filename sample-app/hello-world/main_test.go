package main

import (
	"context"
	"log"
	"net/http"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/golang/mock/gomock"
	"github.com/joho/godotenv"
	"github.com/yuki5155/go-lambda-microkit/domains"
	"github.com/yuki5155/go-lambda-microkit/mocks"
	"github.com/yuki5155/go-lambda-microkit/myaws"
	"github.com/yuki5155/go-lambda-microkit/services"
	"github.com/yuki5155/go-lambda-microkit/utils"
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

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

func getEndPoint() string {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		log.Fatalf("failed to load configuration, %v", err)
	}
	client := cognitoidentityprovider.NewFromConfig(cfg)
	if client == nil {
		log.Fatalf("failed to create Cognito client")
	}
	cloudformationConf := cloudformation.NewFromConfig(cfg)

	cloudFormationClient := myaws.NewCloudFormationClient(cloudformationConf)
	endpoint, err := cloudFormationClient.GetCloudFormationOutput(context.Background(), "sample-app", "HelloWorldAPIwithCUSTOMDOMAIN")
	if err != nil {
		log.Fatalf("Failed to get CloudFormation output: %v", err)
	}

	return endpoint
}

func TestSendRequest(t *testing.T) {
	endPoint := getEndPoint()
	httpUtils := utils.NewHTTPRequestsUtils()
	resp, err := httpUtils.Get(endPoint)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code 200, but got %v", resp.StatusCode)
	}
}
