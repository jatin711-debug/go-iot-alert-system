package service

import (
	pb "alerts/api/proto/alert" // Import generated gRPC code
	"alerts/internal/repository"
	"context"
)

// AlertService handles business logic for alerts
type AlertService struct {
	pb.UnimplementedAlertServiceServer // Embed the generated server interface
	Repo                               *repository.AlertRepository
}

// NewAlertService creates a new instance of AlertService
func NewAlertService(repo *repository.AlertRepository) *AlertService {
	return &AlertService{Repo: repo}
}

// CreateAlert processes and saves an alert
func (s *AlertService) CreateAlert(ctx context.Context, data map[string]interface{}) error {
	return s.Repo.SaveAlert(ctx, data) // Calls repository to store alert in DB
}

// GetAlert fetches an alert by ID
func (s *AlertService) GetAlert(ctx context.Context, id int32) (map[string]interface{}, error) {
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
