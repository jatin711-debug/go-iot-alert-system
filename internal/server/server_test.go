package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	// Create a test Gin router
	gin.SetMode(gin.TestMode)
	router := gin.New()
	SetupRoutes(router, nil) // Setup routes without a handler

	// Create a request to the health endpoint
	req, _ := http.NewRequest("GET", "/api/health", nil)
	w := httptest.NewRecorder()

	// Serve the request
	router.ServeHTTP(w, req)

	// Validate response
	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"status": "ok"}`, w.Body.String())
}
