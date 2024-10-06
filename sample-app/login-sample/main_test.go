package main_test

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/joho/godotenv"
	"github.com/yuki5155/go-lambda-microkit/myaws"
	"github.com/yuki5155/go-lambda-microkit/utils"
)

var (
	apiURL    string
	httpUtils utils.HTTPRequestsUtils
)

func init() {
	log.Println("Initializing test environment...")

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Load AWS configuration
	cfg, err := config.LoadDefaultConfig(context.Background(), config.WithRegion("us-east-1"))
	if err != nil {
		log.Fatalf("Failed to load AWS configuration: %v", err)
	}

	// Create CloudFormation client
	cloudFormationClient := myaws.NewCloudFormationClient(cloudformation.NewFromConfig(cfg))

	// Get UserPoolId from CloudFormation output
	userPoolID, err := cloudFormationClient.GetCloudFormationOutput(context.Background(), "cognito", "UserPoolId")
	if err != nil {
		log.Printf("Failed to get UserPoolId: %v", err)
	} else {
		os.Setenv("COGNITO_USER_POOL_ID", userPoolID)
		log.Printf("Set COGNITO_USER_POOL_ID to: %s", userPoolID)
	}

	// Get UserPoolClientId from CloudFormation output
	userPoolClientID, err := cloudFormationClient.GetCloudFormationOutput(context.Background(), "cognito", "UserPoolClientId")
	if err != nil {
		log.Printf("Failed to get UserPoolClientId: %v", err)
	} else {
		os.Setenv("COGNITO_CLIENT_ID", userPoolClientID)
		log.Printf("Set COGNITO_CLIENT_ID to: %s", userPoolClientID)
	}

	// Get API URL from CloudFormation output
	apiURL, err = cloudFormationClient.GetCloudFormationOutput(context.Background(), "sample-app", "ApiUrl")
	if err != nil {
		log.Printf("Failed to get API URL: %v", err)
	} else {
		log.Printf("API URL set to: %s", apiURL)
	}

	// Verify and log all relevant environment variables
	envVars := []string{"COGNITO_USER_POOL_ID", "COGNITO_CLIENT_ID", "SAMPLEEMAILADDRESS", "SAMPLEPASSWORD"}
	for _, envVar := range envVars {
		value := os.Getenv(envVar)
		if value == "" {
			log.Printf("Warning: %s is not set", envVar)
		} else {
			log.Printf("%s is set (length: %d)", envVar, len(value))
		}
	}

	// Create HTTP request utility
	httpUtils = utils.NewHTTPRequestsUtils()

	log.Println("Test environment initialization completed.")
}

func TestLoginAPI(t *testing.T) {
	// Verify that the required environment variables are set
	requiredEnvVars := []string{"SAMPLEEMAILADDRESS", "SAMPLEPASSWORD", "COGNITO_USER_POOL_ID", "COGNITO_CLIENT_ID"}
	for _, envVar := range requiredEnvVars {
		if os.Getenv(envVar) == "" {
			t.Fatalf("%s is not set", envVar)
		}
	}

	if apiURL == "" {
		t.Fatal("API URL is not set")
	}

	// Construct the login endpoint URL
	loginURL := fmt.Sprintf("%s/login", apiURL)
	t.Logf("Login URL: %s", loginURL)

	// Prepare login request payload
	loginPayload := map[string]string{
		"username": os.Getenv("SAMPLEEMAILADDRESS"),
		"password": os.Getenv("SAMPLEPASSWORD"),
	}
	payloadBytes, err := json.Marshal(loginPayload)
	if err != nil {
		t.Fatalf("Failed to marshal login payload: %v", err)
	}

	// Send POST request to login endpoint
	t.Log("Sending login request...")
	resp, err := httpUtils.Post(loginURL, "application/json", payloadBytes)
	if err != nil {
		t.Fatalf("Failed to send login request: %v", err)
	}

	// Log response details
	t.Logf("Response Status Code: %d", resp.StatusCode)
	t.Logf("Response Body: %s", string(resp.Body))

	// Check response status code
	if resp.StatusCode != 200 {
		t.Errorf("Expected status OK; got %v", resp.StatusCode)
		return
	}

	// Parse response JSON
	var loginResponse struct {
		Token string `json:"token"`
		User  struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		} `json:"user"`
	}
	err = json.Unmarshal(resp.Body, &loginResponse)
	if err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}

	// Validate response
	if loginResponse.Token == "" {
		t.Error("Expected non-empty token in response")
	}
	if loginResponse.User.Name == "" || loginResponse.User.Email == "" {
		t.Error("Expected non-empty user information in response")
	}

	t.Logf("Login successful. User: %s, Email: %s", loginResponse.User.Name, loginResponse.User.Email)
}
