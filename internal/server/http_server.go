package server

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	nrgin "github.com/newrelic/go-agent/v3/integrations/nrgin"
	newrelic "github.com/newrelic/go-agent/v3/newrelic"
)

// SetupNewRelic initializes New Relic agent
func SetupNewRelic(appName, licenseKey string) (*newrelic.Application, error) {
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName(appName),
		newrelic.ConfigLicense(licenseKey),
		newrelic.ConfigDistributedTracerEnabled(true),
	)
	if err != nil {
		return nil, err
	}
	return app, nil
}

// getEnv retrieves the value of an environment variable or returns a default value
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

// SetupRoutes configures all API routes
func SetupRoutes(router *gin.Engine) {
	// Fetch New Relic license key from environment variable or use a default value
	licenseKey := getEnv("NEW_RELIC_LICENSE_KEY", "")
	app, err := SetupNewRelic("alerts", licenseKey)
	if err != nil {
		panic("Failed to initialize New Relic: " + err.Error())
	}
	router.Use(nrgin.Middleware(app)) // New Relic instrumentation
	router.Use(gin.Logger())          // Request logging
	router.Use(gin.Recovery())        // Panic recovery

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
