package telemetry

import (
	"context"

	entities "github.com/nini-k/architecture-sprint-3/microsirvices/telemetry/internal/entities/telemetry"
)

type storage interface {
	// TODO implenent
}

type Service interface {
	ListLatestTelemtryData(ctx context.Context) ([]entities.TelemetryData, error)
	ListTelemtryData(ctx context.Context, pageSize int32, cursor *entities.Cursor) ([]entities.TelemetryData, *entities.Cursor, error)
	PutTelemtryData(ctx context.Context, data entities.TelemetryData) error
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

func (i *Implentation) ListLatestTelemtryData(ctx context.Context) ([]entities.TelemetryData, error) {
	// TODO implement
	return nil, nil
}

func (i *Implentation) ListTelemtryData(ctx context.Context, pageSize int32, cursor *entities.Cursor) ([]entities.TelemetryData, *entities.Cursor, error) {
	// TODO implement
	return nil, nil, nil
}

func (i *Implentation) PutTelemtryData(ctx context.Context, data entities.TelemetryData) error {
	// TODO implement
	return nil
}
