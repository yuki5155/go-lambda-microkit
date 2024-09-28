package myaws_test

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/joho/godotenv"
	"github.com/yuki5155/go-lambda-microkit/myaws"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

func setupCognitoClient(t *testing.T) myaws.CognitoClientInterface {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		t.Fatalf("failed to load configuration, %v", err)
	}
	client := cognitoidentityprovider.NewFromConfig(cfg)
	if client == nil {
		t.Fatalf("failed to create Cognito client")
	}
	cloudformationConf := cloudformation.NewFromConfig(cfg)

	cloudFormationClient := myaws.NewCloudFormationClient(cloudformationConf)
	userPoolID, err := cloudFormationClient.GetCloudFormationOutput(context.Background(), "cognito", "UserPoolId")
	if err != nil {
		t.Fatalf("Failed to get CloudFormation output: %v", err)
	}
	userPoolClientId, err := cloudFormationClient.GetCloudFormationOutput(context.Background(), "cognito", "UserPoolClientId")
	if err != nil {
		t.Fatalf("Failed to get CloudFormation output: %v", err)
	}

	return myaws.NewCognitoClient(client, userPoolClientId, userPoolID)
}

func TestCognitoSignUp(t *testing.T) {
	cognitoClient := setupCognitoClient(t)
	username := os.Getenv("SAMPLEEMAILADDRESS")
	password := os.Getenv("SAMPLEPASSWORD")
	userSub, err := cognitoClient.SignUp(context.Background(), username, password)
	if err != nil {
		t.Errorf("Failed to sign up: %v", err)
	} else {
		t.Logf("UserSub: %s", userSub)
	}
}

func TestCognitoConfirmSignUp(t *testing.T) {
	cognitoClient := setupCognitoClient(t)
	username := os.Getenv("SAMPLEEMAILADDRESS")
	confirmationCode := ""

	err := cognitoClient.ConfirmSignUp(context.Background(), username, confirmationCode)
	if err != nil {
		t.Errorf("Failed to confirm sign up: %v", err)
	} else {
		t.Logf("Successfully confirmed sign up for user: %s", username)
	}
}

func TestCognitoLogin(t *testing.T) {
	cognitoClient := setupCognitoClient(t)
	username := os.Getenv("SAMPLEEMAILADDRESS")
	password := os.Getenv("SAMPLEPASSWORD")

	accessToken, userToke, err := cognitoClient.Login(context.Background(), username, password)
	if userToke == "" {
		t.Fatal("No user token received")
	}
	if err != nil {
		t.Errorf("Failed to login: %v", err)
	} else {
		t.Logf("Successfully logged in. Access Token: %s", accessToken)
	}
	fmt.Println(accessToken)
}

func TestCognitoLogout(t *testing.T) {
	cognitoClient := setupCognitoClient(t)
	username := os.Getenv("SAMPLEEMAILADDRESS")
	password := os.Getenv("SAMPLEPASSWORD")

	// First, login to get an access token
	accessToken, userToken, err := cognitoClient.Login(context.Background(), username, password)
	if err != nil {
		t.Fatalf("Failed to login: %v", err)
	}

	if userToken == "" {
		t.Fatal("No user token received")
	}

	// Now, attempt to logout
	err = cognitoClient.Logout(context.Background(), accessToken)
	if err != nil {
		t.Errorf("Failed to logout: %v", err)
	} else {
		t.Logf("Successfully logged out")
	}

}

func TestCognitoWithAPIGW(t *testing.T) {
	cognitoClient := setupCognitoClient(t)
	username := os.Getenv("SAMPLEEMAILADDRESS")
	password := os.Getenv("SAMPLEPASSWORD")

	// Get the access token using the current implementation
	accessToken, idToken, err := cognitoClient.Login(context.Background(), username, password)
	if err != nil {
		t.Fatalf("Failed to login: %v", err)
	}

	if idToken == "" {
		t.Fatal("No user token received")
	}

	t.Logf("Access Token: %s", accessToken)
	t.Logf("ID Token: %s", idToken)

	if accessToken == "" {
		t.Fatal("No access token received")
	}

	apiURL := os.Getenv("AUTHORIZED_PATH")
	if apiURL == "" {
		t.Fatal("Authorized path not set")
	}
	// Function to make API request
	makeRequest := func(token string) (*http.Response, error) {
		req, err := http.NewRequest("GET", apiURL, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create request: %v", err)
		}
		req.Header.Add("Authorization", "Bearer "+token)

		client := &http.Client{}
		return client.Do(req)
	}

	// Try with Access Token
	t.Log("Trying with Access Token")
	respAccess, err := makeRequest(accessToken)
	if err != nil {
		t.Fatalf("Failed to make API request with Access Token: %v", err)
	}
	defer respAccess.Body.Close()

	bodyAccess, _ := ioutil.ReadAll(respAccess.Body)
	t.Logf("API Response Status (Access Token): %d", respAccess.StatusCode)
	t.Logf("API Response Body (Access Token): %s", string(bodyAccess))

	if respAccess.StatusCode != http.StatusOK {
		t.Errorf("Access Token request failed. Status: %v", respAccess.Status)
	}

	// Note: We can't test with ID Token in the current implementation
	t.Log("Unable to test with ID Token in current implementation")
}
