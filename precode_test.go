package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainHandlerValidRequest(t *testing.T) {
	req, _ := http.NewRequest("GET", "/?city=moscow&count=2", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.NotEmpty(t, responseRecorder.Body.String())
}

func TestMainHandlerInvalidCity(t *testing.T) {
	req, _ := http.NewRequest("GET", "/?city=london&count=2", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	assert.Equal(t, "wrong city value", responseRecorder.Body.String())
}

func TestMainHandlerCountMoreThanTotal(t *testing.T) {
	req, _ := http.NewRequest("GET", "/?city=moscow&count=10", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	expectedAnswer := strings.Join(cafeList["moscow"], ",")
	assert.Equal(t, expectedAnswer, responseRecorder.Body.String())
}
