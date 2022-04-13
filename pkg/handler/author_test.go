package handler

import (
	"bytes"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/magiconair/properties/assert"
	"github.com/stretchr/testify/require"
	"net/http/httptest"
	"rest_api"
	"rest_api/pkg/service"
	mock_service "rest_api/pkg/service/mocks"
	"testing"
)

/*тест переписать
почему-то в параметре структуры появляется перенос строки

*/

func TestHandler_createAuthor (t *testing.T) {

	test := struct {
		name                 string
		inputBody            string
		inputAuthor            rest_api.Author
		expectedStatusCode   int
		expectedResponseBody string
	} {
			name:      "Ok",
			inputBody: `{"authorName": "testAuthor1"}`,
			inputAuthor: rest_api.Author{
				AuthorName:     "testAuthor1",
			},
			expectedStatusCode:   201,
			expectedResponseBody: `{"id":1}`,
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockAuthorService := mock_service.NewMockAuthorService(mockCtrl)
	mockAuthorService.EXPECT().Create(test.inputAuthor).Return(1,nil)
	s := &service.Service{AuthorService: mockAuthorService}
	h := Handler{s}

	r := mux.NewRouter()
	r.HandleFunc("/api/author", h.createAuthor).Methods("POST")

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST","/api/author",bytes.NewBufferString(test.inputBody))


	// Make request
	r.ServeHTTP(w, req)
	// Assert
	assert.Equal(t, w.Code, test.expectedStatusCode)
	require.JSONEq(t, test.expectedResponseBody,w.Body.String())

}