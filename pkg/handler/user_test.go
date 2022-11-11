package handler

import (
	"bytes"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/kirill0909/neurohacking-api/models"
	"github.com/kirill0909/neurohacking-api/pkg/service"
	service_mocks "github.com/kirill0909/neurohacking-api/pkg/service/mocks"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestHandler_signUp(t *testing.T) {
	type mockBehavior func(s *service_mocks.MockUser, user models.User)

	testTable := []struct {
		name                 string
		inputBody            string
		inputUser            models.User
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Ok",
			inputBody: `{"name":"John Down", "email":"john@gmail.com", "password":"JohnPass"}`,
			inputUser: models.User{Name: "John Down", Email: "john@gmail.com", Password: "JohnPass"},
			mockBehavior: func(s *service_mocks.MockUser, user models.User) {
				s.EXPECT().CreateUser(user).Return(1, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"id":1}`,
		},
		{
			name:                 "Invalid body",
			inputBody:            `{"email":"john@gmail.com", "password":"JohnPass"}`,
			inputUser:            models.User{Email: "john@gmail.com", Password: "JohnPass"},
			mockBehavior:         func(s *service_mocks.MockUser, user models.User) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"invalid input body"}`,
		},
		{
			name:                 "Empty valud",
			inputBody:            `{"name":" ", "email":"john@gmail.com", "password":"JohnPass"}`,
			inputUser:            models.User{Email: "john@gmail.com", Password: "JohnPass"},
			mockBehavior:         func(s *service_mocks.MockUser, user models.User) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"invalid input value"}`,
		},
		{
			name:      "Service failure",
			inputBody: `{"name":"John Down", "email":"john@gmail.com", "password":"JohnPass"}`,
			inputUser: models.User{Name: "John Down", Email: "john@gmail.com", Password: "JohnPass"},
			mockBehavior: func(s *service_mocks.MockUser, user models.User) {
				s.EXPECT().CreateUser(user).Return(0, errors.New("something went wrong"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"something went wrong"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			controller := gomock.NewController(t)
			defer controller.Finish()

			user := service_mocks.NewMockUser(controller)
			testCase.mockBehavior(user, testCase.inputUser)

			services := &service.Service{User: user}
			handler := NewHandler(services)

			router := gin.New()
			router.POST("/user/auth/sign-up", handler.signUp)

			recorder := httptest.NewRecorder()
			request := httptest.NewRequest("POST", "/user/auth/sign-up",
				bytes.NewBufferString(testCase.inputBody))

			router.ServeHTTP(recorder, request)

			assert.Equal(t, testCase.expectedResponseBody, recorder.Body.String())
			assert.Equal(t, testCase.expectedStatusCode, recorder.Code)
		})
	}
}

func TestHandler_signIn(t *testing.T) {
	type mockBehavior func(s *service_mocks.MockUser, input models.UserSignInInput)

	testTable := []struct {
		name                 string
		inputBody            string
		inputSignInInput     models.UserSignInInput
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:             "Ok",
			inputBody:        `{"email":"john@gmail.com", "password":"JohnPass"}`,
			inputSignInInput: models.UserSignInInput{Email: "john@gmail.com", Password: "JohnPass"},
			mockBehavior: func(s *service_mocks.MockUser, input models.UserSignInInput) {
				s.EXPECT().GenerateToken(input.Email, input.Password).Return("token", nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"token":"token"}`,
		},
		{
			name:                 "Invalid input boyd",
			inputBody:            `{"password":"JohnPass"}`,
			inputSignInInput:     models.UserSignInInput{Password: "JohnPass"},
			mockBehavior:         func(s *service_mocks.MockUser, input models.UserSignInInput) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"invalid input body"}`,
		},
		{
			name:                 "Invalid input value",
			inputBody:            `{"email":" ", "password":"JohnPass"}`,
			inputSignInInput:     models.UserSignInInput{Password: "JohnPass"},
			mockBehavior:         func(s *service_mocks.MockUser, input models.UserSignInInput) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"invalid input value"}`,
		},
		{
			name:             "Service failure",
			inputBody:        `{"email":"john@gmail.com", "password":"JohnPass"}`,
			inputSignInInput: models.UserSignInInput{Email: "john@gmail.com", Password: "JohnPass"},
			mockBehavior: func(s *service_mocks.MockUser, input models.UserSignInInput) {
				s.EXPECT().GenerateToken(input.Email, input.Password).Return("", errors.New("something went wrong"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"something went wrong"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			controller := gomock.NewController(t)
			defer controller.Finish()

			user := service_mocks.NewMockUser(controller)
			testCase.mockBehavior(user, testCase.inputSignInInput)

			services := &service.Service{User: user}
			handler := NewHandler(services)

			router := gin.New()
			router.POST("/user/auth/sign-in", handler.signIn)

			recorder := httptest.NewRecorder()
			request := httptest.NewRequest("POST", "/user/auth/sign-in",
				bytes.NewBufferString(testCase.inputBody))

			router.ServeHTTP(recorder, request)

			assert.Equal(t, testCase.expectedStatusCode, recorder.Code)
			assert.Equal(t, testCase.expectedResponseBody, recorder.Body.String())
		})
	}

}
