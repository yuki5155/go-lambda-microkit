package services_test

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/yuki5155/go-lambda-microkit/domains"
	"github.com/yuki5155/go-lambda-microkit/mocks"
	"github.com/yuki5155/go-lambda-microkit/services"
)

func TestUserService_Login(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCognitoClient := mocks.NewMockCognitoClientInterface(ctrl)
	userService := services.NewUserService(mockCognitoClient)

	testCases := []struct {
		name          string
		username      string
		password      string
		mockBehavior  func()
		expectedUser  domains.User
		expectedError error
	}{
		{
			name:     "Successful login",
			username: "testuser@example.com",
			password: "password123",
			mockBehavior: func() {
				mockCognitoClient.EXPECT().
					Login(gomock.Any(), "testuser@example.com", "password123").
					Return("", "idtoken123", nil)
			},
			expectedUser: domains.User{
				Name:    "testuser@example.com",
				Email:   "testuser@example.com",
				IDToken: "idtoken123",
			},
			expectedError: nil,
		},
		{
			name:     "Login failure",
			username: "testuser@example.com",
			password: "wrongpassword",
			mockBehavior: func() {
				mockCognitoClient.EXPECT().
					Login(gomock.Any(), "testuser@example.com", "wrongpassword").
					Return("", "", errors.New("login failed"))
			},
			expectedUser:  domains.User{},
			expectedError: errors.New("login failed"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockBehavior()

			user, err := userService.Login(context.Background(), tc.username, tc.password)

			if !reflect.DeepEqual(tc.expectedUser, user) {
				t.Errorf("expected user %v, got %v", tc.expectedUser, user)
			}

			if tc.expectedError == nil && err != nil {
				t.Errorf("expected no error, got %v", err)
			} else if tc.expectedError != nil && err == nil {
				t.Errorf("expected error %v, got nil", tc.expectedError)
			} else if tc.expectedError != nil && err != nil && tc.expectedError.Error() != err.Error() {
				t.Errorf("expected error %v, got %v", tc.expectedError, err)
			}
		})
	}
}
