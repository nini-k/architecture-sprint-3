package device_manager

import (
	"context"
	"errors"
	"fmt"

	entities "github.com/nini-k/architecture-sprint-3/microsirvices/device-manager/internal/entities/device"
	pb "github.com/nini-k/architecture-sprint-3/microsirvices/device-manager/internal/pb/device-manager"
	device_manager "github.com/nini-k/architecture-sprint-3/microsirvices/device-manager/internal/services/device-manager"
	"github.com/nini-k/architecture-sprint-3/microsirvices/device-manager/internal/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ServiceServer struct {
	pb.UnimplementedDeviceManagerServiceServer

	service device_manager.Service
}

func New(
	service device_manager.Service,
) *ServiceServer {
	return &ServiceServer{
		service: service,
	}
}

func (s *ServiceServer) GetDevice(ctx context.Context, req *pb.GetDeviceRequest) (*pb.GetDeviceResponse, error) {
	var (
		id = req.GetId()
	)

	if id <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "id must be greater than zero")
	}

	device, err := s.service.GetDevice(ctx, id)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			return nil, status.Errorf(codes.NotFound, "not found device")
		}

		return nil, fmt.Errorf("service.GetDevice id=%d: %w", id, err)
	}

	return &pb.GetDeviceResponse{
		Device: &pb.Device{
			Name:         device.Name,
			HouseId:      device.HouseID,
			Type:         pb.DeviceType(device.TypeID),
			SerialNumber: device.SerialNumber,
			Status:       device.Status.IsOn,
			CreatedAt:    timestamppb.New(device.CreatedAt),
		},
	}, nil
}

func (s *ServiceServer) AddDevice(ctx context.Context, req *pb.AddDeviceRequest) (*pb.AddDeviceResponse, error) {
	id, err := s.service.AddDevice(ctx, entities.Device{
		Name:         req.GetDevice().GetName(),
		HouseID:      req.GetDevice().GetHouseId(),
		TypeID:       entities.TypeID(req.GetDevice().GetType()),
		SerialNumber: req.GetDevice().GetSerialNumber(),
		Status:       entities.Status{IsOn: req.GetDevice().GetStatus()},
	})
	if err != nil {
		return nil, fmt.Errorf("service.AddDevice: %w", err)
	}

	return &pb.AddDeviceResponse{
		Id: id,
	}, nil
}

func (s *ServiceServer) PutDeviceStatus(ctx context.Context, req *pb.PutDeviceStatusRequest) (*pb.PutDeviceStatusResponse, error) {
	var (
		id = req.GetId()
	)

	if id <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "id must be greater than zero")
	}

	st := entities.Status{IsOn: req.GetStatus()}
	if err := s.service.PutDeviceStatus(ctx, id, st); err != nil {
		return nil, fmt.Errorf("service.PutDeviceStatus id=%d status=%v: %w", id, st, err)
	}

	return &pb.PutDeviceStatusResponse{}, nil
}

func (s *ServiceServer) SetDeviceCommands(ctx context.Context, req *pb.SetDeviceCommandsRequest) (*pb.SetDeviceCommandsResponse, error) {
	var (
		id = req.GetId()
	)

	if id <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "id must be greater than zero")
	}

	var commands entities.Commands
	switch v := req.GetCommands().(type) {
	case *pb.SetDeviceCommandsRequest_Temperature:
		commands = entities.TemperatureCommands{
			TargetTemperature: v.Temperature.GetTargetTemperature(),
		}
	default:
		return nil, status.Errorf(codes.InvalidArgument, "unspecified commands")
	}

	if err := s.service.SetDeviceCommands(ctx, id, commands); err != nil {
		return nil, fmt.Errorf("service.SetDeviceCommands id=%d: %w", id, err)
	}

	return &pb.SetDeviceCommandsResponse{}, nil
}

func (s *ServiceServer) GetDeviceCommands(ctx context.Context, req *pb.GetDeviceCommandsRequest) (*pb.GetDeviceCommandsResponse, error) {
	var (
		id = req.GetId()
	)

	if id <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "id must be greater than zero")
	}

	commands, err := s.service.GetDeviceCommands(ctx, id)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			return nil, status.Errorf(codes.NotFound, "not found device commands")
		}

		return nil, fmt.Errorf("service.GetDeviceCommands: %w", err)
	}

	switch v := commands.(type) {
	case entities.TemperatureCommands:
		return &pb.GetDeviceCommandsResponse{
			Commans: &pb.GetDeviceCommandsResponse_Temperature{
				Temperature: &pb.TemperatureCommands{
					TargetTemperature: v.TargetTemperature,
				},
			},
		}, nil
	default:
		return nil, fmt.Errorf("unspecified commands")
	}
}
