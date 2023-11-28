package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetMovie(t *testing.T) {
	// Create a test router with the same configuration as your main router
	router := gin.Default()
	router.GET("/movies", getMovie)

	// Create a test request
	req, err := http.NewRequest("GET", "/movies?t=Inception&y=2010&type=movie", nil)
	assert.NoError(t, err)

	// Create a test response recorder
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Check the status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Assert the response body, you might need to adjust this based on the actual response from the OMDb API
	expectedResponse := `{"data": "your_expected_response"}`
	assert.JSONEq(t, expectedResponse, w.Body.String())
}
