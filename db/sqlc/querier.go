// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package db

import (
	"context"
)

type Querier interface {
	CreateAlert(ctx context.Context, arg CreateAlertParams) (Alert, error)
	GetAlerts(ctx context.Context, arg GetAlertsParams) ([]Alert, error)
}

var _ Querier = (*Queries)(nil)
