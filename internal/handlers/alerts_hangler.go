package handlers

import (
	"net/http"
	"strconv"

	"alerts/internal/service"

	"github.com/gin-gonic/gin"
)

// Handler wraps the AlertService for dependency injection.
type Handler struct {
	AlertService *service.AlertService
}

// NewHandler creates a new Handler instance.
func NewHandler(alertService *service.AlertService) *Handler {
	return &Handler{AlertService: alertService}
}

// GetAlerts handles HTTP GET requests to fetch an alert by asset ID.
// It expects an "asset_id" query parameter.
func (h *Handler) GetAlerts(c *gin.Context) {
	// Retrieve asset_id from query parameters.
	assetIDStr := c.Query("asset_id")
	if assetIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "asset_id parameter is required"})
		return
	}

	// Convert asset_id to an integer.
	assetID, err := strconv.Atoi(assetIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid asset_id parameter"})
		return
	}

	// Call the AlertService to get the alert by asset ID.
	alert, err := h.AlertService.GetAlert(c.Request.Context(), int32(assetID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Return the alert in the response.
	c.JSON(http.StatusOK, alert)
}
