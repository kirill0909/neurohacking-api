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

func TestHandler_createWord(t *testing.T) {
	type mockBehavior func(s *service_mocks.MockWord, word models.Word, userId, categoryId int)

	testTable := []struct {
		name                 string
		inputBody            string
		inputWord            models.Word
		userId               int
		categoryId           int
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:       "Ok",
			inputBody:  `{"name":"NewWord"}`,
			inputWord:  models.Word{Name: "NewWord"},
			userId:     1,
			categoryId: 1,
			mockBehavior: func(s *service_mocks.MockWord, word models.Word, userId, categoryId int) {
				s.EXPECT().Create(word, userId, categoryId).Return(
					models.Word{Id: 1, UID: 1, CategoryId: 1, Name: "NewWord", DateCreation: "2022-11-13T11:05:47.423733Z", LastUpdate: "2022-11-13T11:05:47.423733Z"}, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"word":{"Id":1,"UID":1,"CategoryId":1,"name":"NewWord","DateCreation":"2022-11-13T11:05:47.423733Z","LastUpdate":"2022-11-13T11:05:47.423733Z"}}`,
		},
		{
			name:                 "Invalid Input Body",
			inputBody:            `{"nam":"NewWord"}`,
			inputWord:            models.Word{},
			userId:               1,
			categoryId:           1,
			mockBehavior:         func(s *service_mocks.MockWord, word models.Word, userId, categoryId int) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"invalid input body"}`,
		},
		{
			name:                 "Invalid Input Value",
			inputBody:            `{"name":" "}`,
			inputWord:            models.Word{},
			userId:               1,
			categoryId:           1,
			mockBehavior:         func(s *service_mocks.MockWord, word models.Word, userId, categoryId int) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"invalid input value"}`,
		},
		{
			name:       "Service failure",
			inputBody:  `{"name":"NewWord"}`,
			inputWord:  models.Word{Name: "NewWord"},
			userId:     1,
			categoryId: 1,
			mockBehavior: func(s *service_mocks.MockWord, word models.Word, userId, categoryId int) {
				s.EXPECT().Create(word, userId, categoryId).Return(models.Word{}, errors.New("something went wrong"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"something went wrong"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			controller := gomock.NewController(t)
			defer controller.Finish()

			word := service_mocks.NewMockWord(controller)
			testCase.mockBehavior(word, testCase.inputWord, testCase.userId, testCase.categoryId)

			service := &service.Service{Word: word}
			handler := NewHandler(service)

			recorder := httptest.NewRecorder()
			request := httptest.NewRequest("POST", "/category/word/",
				bytes.NewBufferString(testCase.inputBody))

			testContext, _ := gin.CreateTestContext(recorder)
			testContext.Request = request
			testContext.Set(userCtx, testCase.userId)
			testContext.AddParam("id", fmt.Sprintf("%d", testCase.categoryId))

			handler.createWord(testContext)

			assert.Equal(t, testCase.expectedStatusCode, recorder.Code)
			assert.Equal(t, testCase.expectedResponseBody, recorder.Body.String())
		})
	}
}

func TestHandler_getAllWords(t *testing.T) {
	type mockBehavior func(s *service_mocks.MockWord, userId, categoryId int)

	testTable := []struct {
		name                 string
		userId               int
		categoryId           int
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedBodyResponse string
	}{
		{
			name:   "Ok",
			userId: 1,
			mockBehavior: func(s *service_mocks.MockWord, userId, categoryId int) {
				s.EXPECT().GetAll(userId, categoryId).Return([]models.Word{
					{Id: 1, UID: 1, CategoryId: 1, Name: "NewWord", DateCreation: "2022-11-12T14:58:21.109514Z", LastUpdate: "2022-11-12T14:58:21.109514Z"},
				}, nil)
			},
			expectedStatusCode:   200,
			expectedBodyResponse: `{"words":[{"Id":1,"UID":1,"CategoryId":1,"name":"NewWord","DateCreation":"2022-11-12T14:58:21.109514Z","LastUpdate":"2022-11-12T14:58:21.109514Z"}]}`,
		},
		{
			name:   "Service Failure",
			userId: 1,
			mockBehavior: func(s *service_mocks.MockWord, userId, categoryId int) {
				s.EXPECT().GetAll(userId, categoryId).Return([]models.Word{}, errors.New("something went wrong"))
			},
			expectedStatusCode:   500,
			expectedBodyResponse: `{"message":"something went wrong"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			controller := gomock.NewController(t)
			controller.Finish()

			word := service_mocks.NewMockWord(controller)
			testCase.mockBehavior(word, testCase.userId, testCase.categoryId)

			service := &service.Service{Word: word}
			handler := NewHandler(service)

			recorder := httptest.NewRecorder()
			request := httptest.NewRequest("GET", "/category/word", nil)

			testContext, _ := gin.CreateTestContext(recorder)
			testContext.Request = request
			testContext.Set(userCtx, 1)
			testContext.AddParam("id", fmt.Sprintf("%d", testCase.categoryId))

			handler.getAllWords(testContext)

			assert.Equal(t, testCase.expectedStatusCode, recorder.Code)
			assert.Equal(t, testCase.expectedBodyResponse, recorder.Body.String())

		})
	}
}

func TestHandler_getById(t *testing.T) {
	type mockBehavior func(s *service_mocks.MockWord, userId, categoryId, wordId int)

	testTable := []struct {
		name                 string
		userId               int
		categoryId           int
		wordId               int
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedBodyResponse string
	}{
		{
			name:       "Ok",
			userId:     1,
			categoryId: 1,
			wordId:     1,
			mockBehavior: func(s *service_mocks.MockWord, userId, categoryId, wordId int) {
				s.EXPECT().GetById(userId, categoryId, wordId).Return(models.Word{
					Id: 1, UID: 1, CategoryId: 1, Name: "NewWord", DateCreation: "2022-11-12T14:58:21.109514Z", LastUpdate: "2022-11-12T14:58:21.109514Z"},
					nil)
			},
			expectedStatusCode:   200,
			expectedBodyResponse: `{"word":{"Id":1,"UID":1,"CategoryId":1,"name":"NewWord","DateCreation":"2022-11-12T14:58:21.109514Z","LastUpdate":"2022-11-12T14:58:21.109514Z"}}`,
		},
		{
			name:       "Service Failure",
			userId:     1,
			categoryId: 1,
			wordId:     1,
			mockBehavior: func(s *service_mocks.MockWord, userId, categoryId, wordId int) {
				s.EXPECT().GetById(userId, categoryId, wordId).Return(models.Word{}, errors.New("something went wrong"))
			},
			expectedStatusCode:   500,
			expectedBodyResponse: `{"message":"something went wrong"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			controller := gomock.NewController(t)
			controller.Finish()

			word := service_mocks.NewMockWord(controller)
			testCase.mockBehavior(word, testCase.userId, testCase.categoryId, testCase.wordId)

			service := &service.Service{Word: word}
			handler := NewHandler(service)

			recorder := httptest.NewRecorder()
			request := httptest.NewRequest("GET", "/category/word/", nil)

			testContext, _ := gin.CreateTestContext(recorder)
			testContext.Request = request
			testContext.Set(userCtx, 1)
			testContext.AddParam("id", fmt.Sprintf("%d", testCase.categoryId))
			testContext.AddParam("word_id", fmt.Sprintf("%d", testCase.wordId))

			handler.getWordById(testContext)

			assert.Equal(t, testCase.expectedStatusCode, recorder.Code)
			assert.Equal(t, testCase.expectedBodyResponse, recorder.Body.String())

		})
	}
}

func TestHandler_updateWord(t *testing.T) {
	type mockBehavior func(s *service_mocks.MockWord, word models.WordUpdateInput,
		userId, categoryId, wordId int)

	testTable := []struct {
		name                 string
		inputBody            string
		inputCategory        models.WordUpdateInput
		userId               int
		categoryId           int
		wordId               int
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedBodyResponse string
	}{
		{
			name:          "Ok",
			inputBody:     `{"name":"NewWord"}`,
			inputCategory: models.WordUpdateInput{Name: "NewWord"},
			userId:        1,
			categoryId:    1,
			wordId:        1,
			mockBehavior: func(s *service_mocks.MockWord, word models.WordUpdateInput, userId, categoryId, wordId int) {
				s.EXPECT().Update(word, userId, categoryId, wordId).Return(models.Word{Id: 1, UID: 1, CategoryId: 1, Name: "NewWord",
					DateCreation: "2022-11-12T14:58:21.109514Z", LastUpdate: "2022-11-12T14:58:21.109514Z"}, nil)
			},
			expectedStatusCode:   200,
			expectedBodyResponse: `{"word":{"Id":1,"UID":1,"CategoryId":1,"name":"NewWord","DateCreation":"2022-11-12T14:58:21.109514Z","LastUpdate":"2022-11-12T14:58:21.109514Z"}}`,
		},
		{
			name:                 "Invalid Input Body",
			inputBody:            `{"nam":"NewWord"}`,
			inputCategory:        models.WordUpdateInput{},
			userId:               1,
			categoryId:           1,
			wordId:               1,
			mockBehavior:         func(s *service_mocks.MockWord, word models.WordUpdateInput, userId, categoryId, wordId int) {},
			expectedStatusCode:   400,
			expectedBodyResponse: `{"message":"invalid input body"}`,
		},
		{
			name:                 "Invalid Input Value",
			inputBody:            `{"name":" "}`,
			inputCategory:        models.WordUpdateInput{},
			userId:               1,
			categoryId:           1,
			wordId:               1,
			mockBehavior:         func(s *service_mocks.MockWord, word models.WordUpdateInput, userId, categoryId, wordId int) {},
			expectedStatusCode:   400,
			expectedBodyResponse: `{"message":"invalid input value"}`,
		},
		{
			name:          "Service Failure",
			inputBody:     `{"name":"NewWord"}`,
			inputCategory: models.WordUpdateInput{Name: "NewWord"},
			userId:        1,
			categoryId:    1,
			wordId:        1,
			mockBehavior: func(s *service_mocks.MockWord, word models.WordUpdateInput, userId, categoryId, wordId int) {
				s.EXPECT().Update(word, userId, categoryId, wordId).Return(models.Word{}, errors.New("something went wrong"))
			},
			expectedStatusCode:   500,
			expectedBodyResponse: `{"message":"something went wrong"}`,
		}}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			controller := gomock.NewController(t)
			controller.Finish()

			word := service_mocks.NewMockWord(controller)
			testCase.mockBehavior(word, testCase.inputCategory,
				testCase.userId, testCase.categoryId, testCase.wordId)

			service := &service.Service{Word: word}
			handler := NewHandler(service)

			recorder := httptest.NewRecorder()
			request := httptest.NewRequest("PUT", "/category/word",
				bytes.NewBufferString(testCase.inputBody))

			testContext, _ := gin.CreateTestContext(recorder)
			testContext.Request = request
			testContext.Set(userCtx, 1)
			testContext.AddParam("id", fmt.Sprintf("%d", testCase.categoryId))
			testContext.AddParam("word_id", fmt.Sprintf("%d", testCase.wordId))

			handler.updateWord(testContext)

			assert.Equal(t, testCase.expectedStatusCode, recorder.Code)
			assert.Equal(t, testCase.expectedBodyResponse, recorder.Body.String())

		})
	}
}

func TestHandler_deleteWord(t *testing.T) {
	type mockBehavior func(s *service_mocks.MockWord, userId, categoryId, wordId int)

	testTable := []struct {
		name                 string
		userId               int
		categoryId           int
		wordId               int
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedBodyResponse string
	}{
		{
			name:       "Ok",
			userId:     1,
			categoryId: 1,
			wordId:     1,
			mockBehavior: func(s *service_mocks.MockWord, userId, categoryId, wordId int) {
				s.EXPECT().Delete(userId, categoryId, wordId).Return(models.Word{
					Id: 1, UID: 1, CategoryId: 1, Name: "DeletedWord", DateCreation: "2022-11-12T14:58:21.109514Z", LastUpdate: "2022-11-12T14:58:21.109514Z"},
					nil)
			},
			expectedStatusCode:   200,
			expectedBodyResponse: `{"word":{"Id":1,"UID":1,"CategoryId":1,"name":"DeletedWord","DateCreation":"2022-11-12T14:58:21.109514Z","LastUpdate":"2022-11-12T14:58:21.109514Z"}}`,
		},
		{
			name:       "Service Failure",
			userId:     1,
			categoryId: 1,
			wordId:     1,
			mockBehavior: func(s *service_mocks.MockWord, userId, categoryId, wordId int) {
				s.EXPECT().Delete(userId, categoryId, wordId).Return(models.Word{}, errors.New("something went wrong"))
			},
			expectedStatusCode:   500,
			expectedBodyResponse: `{"message":"something went wrong"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			controller := gomock.NewController(t)
			controller.Finish()

			word := service_mocks.NewMockWord(controller)
			testCase.mockBehavior(word, testCase.userId, testCase.categoryId, testCase.wordId)

			service := &service.Service{Word: word}
			handler := NewHandler(service)

			recorder := httptest.NewRecorder()
			request := httptest.NewRequest("DELETE", "/category/word/", nil)

			testContext, _ := gin.CreateTestContext(recorder)
			testContext.Request = request
			testContext.Set(userCtx, 1)
			testContext.AddParam("id", fmt.Sprintf("%d", testCase.categoryId))
			testContext.AddParam("word_id", fmt.Sprintf("%d", testCase.wordId))

			handler.deleteWord(testContext)

			assert.Equal(t, testCase.expectedStatusCode, recorder.Code)
			assert.Equal(t, testCase.expectedBodyResponse, recorder.Body.String())

		})
	}
}
