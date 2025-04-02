package server

import (
	"net/http"
	// "os"

	"github.com/gin-gonic/gin"
	nrgin "github.com/newrelic/go-agent/v3/integrations/nrgin"
	newrelic "github.com/newrelic/go-agent/v3/newrelic"
	route_handler "alerts/internal/handlers"
	utils "alerts/internal/utils"
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

// SetupRoutes configures all API routes
func SetupRoutes(router *gin.Engine) {
	// Fetch New Relic license key from environment variable or use a default value
	licenseKey := utils.GetEnv("NEW_RELIC_LICENSE_KEY", "")
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

		api.GET("/alerts", route_handler.GetAlerts)   // Get all alerts
	}

}

// HealthCheck responds with a simple status
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
