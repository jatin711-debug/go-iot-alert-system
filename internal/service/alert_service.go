package service

import (
	pb "alerts/api/proto/alert" // Import generated gRPC code
	"alerts/internal/cache"
	"alerts/internal/repository"
	"context"
	"fmt"
	"time"
)

// AlertService handles business logic for alerts
type AlertService struct {
	pb.UnimplementedAlertServiceServer
	Repo         *repository.AlertRepository
	CacheManager *cache.CacheManager
}

// NewAlertService creates a new instance of AlertService
func NewAlertService(repo *repository.AlertRepository, cacheManager *cache.CacheManager) *AlertService {
	return &AlertService{Repo: repo, CacheManager: cacheManager}
}

// CreateAlert processes and saves an alert
func (s *AlertService) CreateAlert(ctx context.Context, data map[string]any) error {
	// Check if alert already exists in cache, update it or else create a new one
	if _, err := s.CacheManager.Get(fmt.Sprintf("alert_%d", data["asset_id"])); err == nil {
		// Alert exists in cache, update it
		if err := s.CacheManager.Set(fmt.Sprintf("alert_%d", data["asset_id"]), data, 30*time.Second); err != nil {
			return fmt.Errorf("failed to update alert in cache: %w", err)
		}
	} else {
		// Alert does not exist, create a new one
		if err := s.CacheManager.Set(fmt.Sprintf("alert_%d", data["asset_id"]), data, 30*time.Second); err != nil {
			return fmt.Errorf("failed to create alert in cache: %w", err)
		}
	}

	return s.Repo.SaveAlert(ctx, data) // Calls repository to store alert in DB
}

// GetAlert fetches an alert by ID
func (s *AlertService) GetAlert(ctx context.Context, id int32) (map[string]any, error) {

	// Check cache first with a key alert_id
	if alert, err := s.CacheManager.Get(fmt.Sprintf("alert_%d", id)); err == nil {
		return alert, nil // Return cached alert if found
	}
	return s.Repo.FindAlertByID(ctx, id)
}

func (s *AlertService) GetAlerts(ctx context.Context, req *pb.AlertRequest) (*pb.AlertResponse, error) {
	alerts, err := s.Repo.FindAlertByID(ctx, req.AssetId)
	if err != nil {
		return nil, err
	}
	response := &pb.AlertResponse{
		Alerts: make([]*pb.Alert, len(alerts)),
	}
	return response, nil
}
