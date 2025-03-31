package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configures all API routes
func SetupRoutes(router *gin.Engine) {
	// Middleware
	router.Use(gin.Logger())   // Request logging
	router.Use(gin.Recovery()) // Panic recovery

	// API Routes
	api := router.Group("/api")
	{
		api.GET("/health", HealthCheck) // Health Check
	}
}

// HealthCheck responds with a simple status
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
