package repository

import (
	sqlc "alerts/db/sqlc"
	"context"
	"fmt"
)

// AlertRepository defines methods for alert data access
type AlertRepository struct {
	Db *sqlc.Queries // This is the SQLC-generated DB queries interface
}

// NewAlertRepository creates a new instance of AlertRepository
func NewAlertRepository(db *sqlc.Queries) *AlertRepository {
	return &AlertRepository{Db: db}
}

// SaveAlert saves a new alert to the database
func (r *AlertRepository) SaveAlert(ctx context.Context, data map[string]any) error {
	// Example of a query from SQLC
	_, err := r.Db.CreateAlert(ctx, sqlc.CreateAlertParams{
		AssetID:  data["asset_id"].(int32),
		Severity: data["severity"].(string),
	})

	return err
}

// FindAlertByID fetches an alert from the database by ID
func (r *AlertRepository) FindAlertByID(ctx context.Context, id int32) (map[string]any, error) {
	alert, err := r.Db.GetAlerts(ctx, sqlc.GetAlertsParams{
		AssetID:  id,
		Severity: "High",
	})
	if err != nil {
		return nil, err
	}

	if len(alert) == 0 {
		return nil, fmt.Errorf("alert not found")
	}

	return map[string]any{
		"id":         alert[0].ID,
		"asset_id":   alert[0].AssetID,
		"severity":   alert[0].Severity,
		"created_at": alert[0].CreatedAt,
	}, nil
}
