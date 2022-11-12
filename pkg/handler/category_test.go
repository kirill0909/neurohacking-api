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

func TestHandler_createCategory(t *testing.T) {
	type mockBehavior func(s *service_mocks.MockCategory, category models.Category, userId int)

	testTable := []struct {
		name                 string
		inputBody            string
		inputCategory        models.Category
		userId               int
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedBodyResponse string
	}{
		{
			name:          "Ok",
			inputBody:     `{"name":"NewCategory"}`,
			inputCategory: models.Category{Name: "NewCategory"},
			userId:        1,
			mockBehavior: func(s *service_mocks.MockCategory, category models.Category, userId int) {
				s.EXPECT().Create(category, userId).Return(models.Category{Id: 1, UID: 1, Name: "NewCategory",
					DateCreation: "2022-11-12T14:58:21.109514Z", LastUpdate: "2022-11-12T14:58:21.109514Z"}, nil)
			},
			expectedStatusCode:   200,
			expectedBodyResponse: `{"category":{"Id":1,"UID":1,"name":"NewCategory","DateCreation":"2022-11-12T14:58:21.109514Z","LastUpdate":"2022-11-12T14:58:21.109514Z"}}`,
		},
		{
			name:                 "Invalid Input Body",
			inputBody:            `{"nam":"NewCategory"}`,
			inputCategory:        models.Category{Name: "NewCategory"},
			userId:               1,
			mockBehavior:         func(s *service_mocks.MockCategory, category models.Category, userId int) {},
			expectedStatusCode:   400,
			expectedBodyResponse: `{"message":"invalid input body"}`,
		},
		{
			name:                 "Invalid Input Value",
			inputBody:            `{"name":" "}`,
			inputCategory:        models.Category{Name: "NewCategory"},
			userId:               1,
			mockBehavior:         func(s *service_mocks.MockCategory, category models.Category, userId int) {},
			expectedStatusCode:   400,
			expectedBodyResponse: `{"message":"invalid input value"}`,
		},
		{
			name:          "Service Failure",
			inputBody:     `{"name":"NewCategory"}`,
			inputCategory: models.Category{Name: "NewCategory"},
			userId:        1,
			mockBehavior: func(s *service_mocks.MockCategory, category models.Category, userId int) {
				s.EXPECT().Create(category, userId).Return(models.Category{}, errors.New("something went wrong"))
			},
			expectedStatusCode:   500,
			expectedBodyResponse: `{"message":"something went wrong"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			controller := gomock.NewController(t)
			controller.Finish()

			category := service_mocks.NewMockCategory(controller)
			testCase.mockBehavior(category, testCase.inputCategory, testCase.userId)

			service := &service.Service{Category: category}
			handler := NewHandler(service)

			recorder := httptest.NewRecorder()
			request := httptest.NewRequest("POST", "/category",
				bytes.NewBufferString(testCase.inputBody))

			testContext, _ := gin.CreateTestContext(recorder)
			testContext.Request = request
			testContext.Set(userCtx, 1)

			handler.createCategory(testContext)

			assert.Equal(t, testCase.expectedStatusCode, recorder.Code)
			assert.Equal(t, testCase.expectedBodyResponse, recorder.Body.String())

		})
	}
}
