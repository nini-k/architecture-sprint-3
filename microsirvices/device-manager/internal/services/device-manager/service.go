package device_manager

import (
	"context"
	"fmt"

	entities "github.com/nini-k/architecture-sprint-3/microsirvices/device-manager/internal/entities/device"
)

type storage interface {
	GetDevice(ctx context.Context, id int64) (entities.Device, error)
	AddDevice(ctx context.Context, device entities.Device) (int64, error)
	PutDeviceStatus(ctx context.Context, id int64, status entities.Status) error
	SetDeviceCommands(ctx context.Context, id int64, commands entities.Commands) error
	GetDeviceCommands(ctx context.Context, id int64) (entities.Commands, error)
}

type Service interface {
	GetDevice(ctx context.Context, id int64) (entities.Device, error)
	AddDevice(ctx context.Context, device entities.Device) (int64, error)
	PutDeviceStatus(ctx context.Context, id int64, status entities.Status) error
	SetDeviceCommands(ctx context.Context, id int64, commands entities.Commands) error
	GetDeviceCommands(ctx context.Context, id int64) (entities.Commands, error)
}

type Implentation struct {
	storage storage
}

func New(
	storage storage,
) *Implentation {
	return &Implentation{
		storage: storage,
	}
}

func (i *Implentation) GetDevice(ctx context.Context, id int64) (entities.Device, error) {
	device, err := i.storage.GetDevice(ctx, id)
	if err != nil {
		return entities.Device{}, fmt.Errorf("storage.GetDevice id=%d: %w", id, err)
	}

	return device, nil
}

func (i *Implentation) AddDevice(ctx context.Context, device entities.Device) (int64, error) {
	id, err := i.storage.AddDevice(ctx, device)
	if err != nil {
		return 0, fmt.Errorf("storage.AddDevice: %w", err)
	}

	return id, nil
}

func (i *Implentation) PutDeviceStatus(ctx context.Context, id int64, status entities.Status) error {
	if err := i.storage.PutDeviceStatus(ctx, id, status); err != nil {
		return fmt.Errorf("storage.PutDeviceStatus id=%d status=%v: %w", id, status, err)
	}

	return nil
}

func (i *Implentation) SetDeviceCommands(ctx context.Context, id int64, commands entities.Commands) error {
	if err := i.storage.SetDeviceCommands(ctx, id, commands); err != nil {
		return fmt.Errorf("SetDeviceCommands: %w", err)
	}

	return nil
}

func (i *Implentation) GetDeviceCommands(ctx context.Context, id int64) (entities.Commands, error) {
	commands, err := i.storage.GetDeviceCommands(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("storage.GetDeviceCommands: %w", err)
	}

	return commands, nil
}
