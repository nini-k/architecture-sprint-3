package telemetry

import (
	"context"

	pb "github.com/nini-k/architecture-sprint-3/microsirvices/telemetry/internal/pb/telemetry"
	telemetry "github.com/nini-k/architecture-sprint-3/microsirvices/telemetry/internal/services/telemetry"
)

type ServiceServer struct {
	pb.UnimplementedTelemetryServiceServer

	service telemetry.Service
}

func New(
	service telemetry.Service,
) *ServiceServer {
	return &ServiceServer{
		service: service,
	}
}

func (s *ServiceServer) ListLatestDeviceTelemetry(ctx context.Context, req *pb.ListLatestDeviceTelemetryRequest) (*pb.ListLatestDeviceTelemetryResponse, error) {
	// TODO implement
	return nil, nil
}

func (s *ServiceServer) ListDeviceTelemetry(ctx context.Context, req *pb.ListDeviceTelemetryRequest) (*pb.ListDeviceTelemetryResponse, error) {
	// TODO implement
	return nil, nil
}
