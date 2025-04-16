package server

import (
	"net/http"

	"alerts/internal/handlers"
	 utils "alerts/internal/utils"

	"github.com/gin-gonic/gin"
	nrgin "github.com/newrelic/go-agent/v3/integrations/nrgin"
	newrelic "github.com/newrelic/go-agent/v3/newrelic"
)

// SetupNewRelic initializes the New Relic agent.
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

// SetupRoutes configures all API routes.
// Note: We now accept a handler instance for dependency injection.
func SetupRoutes(router *gin.Engine, alertHandler *handlers.Handler) {
	// Set Gin to release mode for production.
	// This is important for performance and security in production environments.
	gin.SetMode(gin.ReleaseMode)
	// Fetch New Relic license key from environment variables.
	licenseKey := utils.GetEnv("NEW_RELIC_LICENSE_KEY", "")
	app, err := SetupNewRelic("alerts", licenseKey)
	if err != nil {
		panic("Failed to initialize New Relic: " + err.Error())
	}
	router.Use(nrgin.Middleware(app)) // New Relic instrumentation.
	router.Use(gin.Logger())          // Request logging.
	router.Use(gin.Recovery())        // Panic recovery.

	// API Routes.
	api := router.Group("/api")
	{
		api.GET("/health", HealthCheck)         // Health Check.
		api.GET("/alerts", alertHandler.GetAlerts) // Get alert by asset_id.
	}
}

// HealthCheck responds with a simple status.
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
