package server

import (
	"backend/internal/product"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServer(test *testing.T) {

	//GIVEN
	testServer, err := New(Config{
		Port: 8080,
	})
	assert.NoError(test, err, "no error should happened when building the request")
	recorder := httptest.NewRecorder()

	inputProduct := product.Product{
		Name: "Test Product",
	}
	inputProductJSON, err := json.Marshal(inputProduct)
	assert.NoError(test, err, "should be able to marshal product")

	req, err := http.NewRequest("POST", "/admin/products", bytes.NewReader(inputProductJSON))
	assert.NoError(test, err, "no error should happened when building the request")
	req.Header.Add("Authorization", "ABC")

	//WHEN
	testServer.Engine.ServeHTTP(recorder, req)

	//THEN
	assert.Equal(test, http.StatusOK, recorder.Code)
}
