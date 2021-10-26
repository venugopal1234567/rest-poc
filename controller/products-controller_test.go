package controller

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"rest-poc/controller/servicemock"
	"testing"
)

func TestGeAllProducts(t *testing.T) {

	tests := []struct {
		name             string
		serviceMock      *servicemock.ServiceMock
		expectedRespCode int
	}{
		{
			"200 Ok",
			&servicemock.ServiceMock{ErrResp: nil},
			http.StatusOK,
		},
		{
			"500 internal server error",
			&servicemock.ServiceMock{ErrResp: errors.New("failed to get products")},
			http.StatusInternalServerError,
		},
	}

	for _, test := range tests {
		fmt.Println(test.name)
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/products", nil)

		c := NewProductController(test.serviceMock)
		c.GeAllProducts(w, r)

		if w.Code != test.expectedRespCode {
			t.Errorf("Expected status: %d, but got: %d", test.expectedRespCode, w.Code)
		}
	}
}

func TestAddProduct(t *testing.T) {

	const product = `{
		"id": "kpln1233",
		"name": "apple",
		"cost": 55,
		"quantity": 33
	}`
	tests := []struct {
		name             string
		serviceMock      *servicemock.ServiceMock
		reqBody          []byte
		expectedRespCode int
	}{
		{
			"200 Ok",
			&servicemock.ServiceMock{ErrResp: nil},
			[]byte(product),
			http.StatusOK,
		},
		{
			"400  Bad Request",
			&servicemock.ServiceMock{ErrResp: nil},
			[]byte("bjbjbnbmnn"),
			http.StatusBadRequest,
		},
		{
			"500 Internal error",
			&servicemock.ServiceMock{ErrResp: errors.New("failed to add product")},
			[]byte(product),
			http.StatusInternalServerError,
		},
	}

	for _, test := range tests {
		fmt.Println(test.name)
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(test.reqBody))
		c := NewProductController(test.serviceMock)
		c.AddProduct(w, r)

		if w.Code != test.expectedRespCode {
			t.Errorf("Expected status: %d, but got: %d", test.expectedRespCode, w.Code)
		}
	}
}
