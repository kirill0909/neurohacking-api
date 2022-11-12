package handler

import (
	"bytes"
	"errors"
	"fmt"
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

func TestHandler_getAllCategories(t *testing.T) {
	type mockBehavior func(s *service_mocks.MockCategory, userId int)

	testTable := []struct {
		name                 string
		userId               int
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedBodyResponse string
	}{
		{
			name:   "Ok",
			userId: 1,
			mockBehavior: func(s *service_mocks.MockCategory, userId int) {
				s.EXPECT().GetAll(userId).Return([]models.Category{
					{Id: 1, UID: 1, Name: "NewCategory", DateCreation: "2022-11-12T14:58:21.109514Z", LastUpdate: "2022-11-12T14:58:21.109514Z"},
				}, nil)
			},
			expectedStatusCode:   200,
			expectedBodyResponse: `{"categories":[{"Id":1,"UID":1,"name":"NewCategory","DateCreation":"2022-11-12T14:58:21.109514Z","LastUpdate":"2022-11-12T14:58:21.109514Z"}]}`,
		},
		{
			name:   "Service Failure",
			userId: 1,
			mockBehavior: func(s *service_mocks.MockCategory, userId int) {
				s.EXPECT().GetAll(userId).Return([]models.Category{}, errors.New("something went wrong"))
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
			testCase.mockBehavior(category, testCase.userId)

			service := &service.Service{Category: category}
			handler := NewHandler(service)

			recorder := httptest.NewRecorder()
			request := httptest.NewRequest("GET", "/category", nil)

			testContext, _ := gin.CreateTestContext(recorder)
			testContext.Request = request
			testContext.Set(userCtx, 1)

			handler.getAllCategories(testContext)

			assert.Equal(t, testCase.expectedStatusCode, recorder.Code)
			assert.Equal(t, testCase.expectedBodyResponse, recorder.Body.String())

		})
	}
}

func TestHandler_getCategoryById(t *testing.T) {
	type mockBehavior func(s *service_mocks.MockCategory, userId, categoryId int)

	testTable := []struct {
		name                 string
		userId               int
		categoryId           int
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedBodyResponse string
	}{
		{
			name:       "Ok",
			userId:     1,
			categoryId: 1,
			mockBehavior: func(s *service_mocks.MockCategory, userId, categoryId int) {
				s.EXPECT().GetById(userId, categoryId).Return(models.Category{
					Id: 1, UID: 1, Name: "NewCategory", DateCreation: "2022-11-12T14:58:21.109514Z", LastUpdate: "2022-11-12T14:58:21.109514Z"},
					nil)
			},
			expectedStatusCode:   200,
			expectedBodyResponse: `{"category":{"Id":1,"UID":1,"name":"NewCategory","DateCreation":"2022-11-12T14:58:21.109514Z","LastUpdate":"2022-11-12T14:58:21.109514Z"}}`,
		},
		{
			name:       "Service Failure",
			userId:     1,
			categoryId: 1,
			mockBehavior: func(s *service_mocks.MockCategory, userId, categoryId int) {
				s.EXPECT().GetById(userId, categoryId).Return(models.Category{}, errors.New("something went wrong"))
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
			testCase.mockBehavior(category, testCase.userId, testCase.categoryId)

			service := &service.Service{Category: category}
			handler := NewHandler(service)

			recorder := httptest.NewRecorder()
			request := httptest.NewRequest("GET", "/category/", nil)

			testContext, _ := gin.CreateTestContext(recorder)
			testContext.Request = request
			testContext.Set(userCtx, 1)
			testContext.AddParam("id", fmt.Sprintf("%d", testCase.categoryId))

			handler.getCategoryById(testContext)

			assert.Equal(t, testCase.expectedStatusCode, recorder.Code)
			assert.Equal(t, testCase.expectedBodyResponse, recorder.Body.String())

		})
	}
}

func TestHandler_updateCategory(t *testing.T) {
	type mockBehavior func(s *service_mocks.MockCategory, category models.CategoryUpdateInput,
		userId, categoryId int)

	testTable := []struct {
		name                 string
		inputBody            string
		inputCategory        models.CategoryUpdateInput
		userId               int
		categoryId           int
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedBodyResponse string
	}{
		{
			name:          "Ok",
			inputBody:     `{"name":"NewCategory"}`,
			inputCategory: models.CategoryUpdateInput{Name: stringToPointer("NewCategory")},
			userId:        1,
			categoryId:    1,
			mockBehavior: func(s *service_mocks.MockCategory, category models.CategoryUpdateInput, userId, categoryId int) {
				s.EXPECT().Update(category, userId, categoryId).Return(models.Category{Id: 1, UID: 1, Name: "NewCategory",
					DateCreation: "2022-11-12T14:58:21.109514Z", LastUpdate: "2022-11-12T14:58:21.109514Z"}, nil)
			},
			expectedStatusCode:   200,
			expectedBodyResponse: `{"category":{"Id":1,"UID":1,"name":"NewCategory","DateCreation":"2022-11-12T14:58:21.109514Z","LastUpdate":"2022-11-12T14:58:21.109514Z"}}`,
		},
		{
			name:                 "Invalid Input Body",
			inputBody:            `{"nam":"NewCategory"}`,
			inputCategory:        models.CategoryUpdateInput{},
			userId:               1,
			categoryId:           1,
			mockBehavior:         func(s *service_mocks.MockCategory, category models.CategoryUpdateInput, userId, categoryId int) {},
			expectedStatusCode:   400,
			expectedBodyResponse: `{"message":"invalid input body"}`,
		},
		{
			name:                 "Invalid Input Value",
			inputBody:            `{"name":" "}`,
			inputCategory:        models.CategoryUpdateInput{},
			userId:               1,
			categoryId:           1,
			mockBehavior:         func(s *service_mocks.MockCategory, category models.CategoryUpdateInput, userId, categoryId int) {},
			expectedStatusCode:   400,
			expectedBodyResponse: `{"message":"invalid input value"}`,
		},
		{
			name:          "Service Failuer",
			inputBody:     `{"name":"NewCategory"}`,
			inputCategory: models.CategoryUpdateInput{Name: stringToPointer("NewCategory")},
			userId:        1,
			categoryId:    1,
			mockBehavior: func(s *service_mocks.MockCategory, category models.CategoryUpdateInput, userId, categoryId int) {
				s.EXPECT().Update(category, userId, categoryId).Return(models.Category{}, errors.New("something went wrong"))
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
			testCase.mockBehavior(category, testCase.inputCategory,
				testCase.userId, testCase.categoryId)

			service := &service.Service{Category: category}
			handler := NewHandler(service)

			recorder := httptest.NewRecorder()
			request := httptest.NewRequest("PUT", "/category",
				bytes.NewBufferString(testCase.inputBody))

			testContext, _ := gin.CreateTestContext(recorder)
			testContext.Request = request
			testContext.Set(userCtx, 1)
			testContext.AddParam("id", fmt.Sprintf("%d", testCase.categoryId))

			handler.updateCategory(testContext)

			assert.Equal(t, testCase.expectedStatusCode, recorder.Code)
			assert.Equal(t, testCase.expectedBodyResponse, recorder.Body.String())

		})
	}
}

func TestHandler_deleteCategory(t *testing.T) {
	type mockBehavior func(s *service_mocks.MockCategory, userId, categoryId int)

	testTable := []struct {
		name                 string
		userId               int
		categoryId           int
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedBodyResponse string
	}{
		{
			name:       "Ok",
			userId:     1,
			categoryId: 1,
			mockBehavior: func(s *service_mocks.MockCategory, userId, categoryId int) {
				s.EXPECT().Delete(userId, categoryId).Return(models.Category{
					Id: 1, UID: 1, Name: "DeletedCategory", DateCreation: "2022-11-12T14:58:21.109514Z", LastUpdate: "2022-11-12T14:58:21.109514Z"},
					nil)
			},
			expectedStatusCode:   200,
			expectedBodyResponse: `{"category":{"Id":1,"UID":1,"name":"DeletedCategory","DateCreation":"2022-11-12T14:58:21.109514Z","LastUpdate":"2022-11-12T14:58:21.109514Z"}}`,
		},
		{
			name:       "Service Failure",
			userId:     1,
			categoryId: 1,
			mockBehavior: func(s *service_mocks.MockCategory, userId, categoryId int) {
				s.EXPECT().Delete(userId, categoryId).Return(models.Category{}, errors.New("something went wrong"))
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
			testCase.mockBehavior(category, testCase.userId, testCase.categoryId)

			service := &service.Service{Category: category}
			handler := NewHandler(service)

			recorder := httptest.NewRecorder()
			request := httptest.NewRequest("DELETE", "/category/", nil)

			testContext, _ := gin.CreateTestContext(recorder)
			testContext.Request = request
			testContext.Set(userCtx, 1)
			testContext.AddParam("id", fmt.Sprintf("%d", testCase.categoryId))

			handler.deleteCategory(testContext)

			assert.Equal(t, testCase.expectedStatusCode, recorder.Code)
			assert.Equal(t, testCase.expectedBodyResponse, recorder.Body.String())

		})
	}
}
