package device

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	entities "github.com/nini-k/architecture-sprint-3/microsirvices/device-manager/internal/entities/device"
	"github.com/nini-k/architecture-sprint-3/microsirvices/device-manager/internal/storage"
)

var (
	psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
)

type Storage struct {
	pool *pgxpool.Pool
}

func New(
	pool *pgxpool.Pool,
) *Storage {
	return &Storage{
		pool: pool,
	}
}

func (s *Storage) GetDevice(ctx context.Context, id int64) (entities.Device, error) {
	sql, args, err := psql.Select(
		"id",
		"house_id",
		"type_id",
		"serial_number",
		"status",
		"created_at",
	).From("device").Where("id = ?", id).ToSql()
	if err != nil {
		return entities.Device{}, fmt.Errorf("psql.ToSql: %w", err)
	}

	row := s.pool.QueryRow(ctx, sql, args...)

	var out entities.Device
	if err := row.Scan(
		&out.ID,
		&out.HouseID,
		&out.TypeID,
		&out.SerialNumber,
		&out.Status.IsOn,
		&out.CreatedAt,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entities.Device{}, storage.ErrNotFound
		}

		return entities.Device{}, fmt.Errorf("row.Scan: %w", err)
	}

	return out, nil
}

func (s *Storage) AddDevice(ctx context.Context, device entities.Device) (int64, error) {
	sql, args, err := psql.Insert("device").Columns(
		"name",
		"house_id",
		"type_id",
		"serial_number",
		"status",
	).Values(device.Name, device.HouseID, device.TypeID, device.SerialNumber, device.Status.IsOn).
		Suffix("returning id").ToSql()
	if err != nil {
		return 0, fmt.Errorf("psql.ToSql: %w", err)
	}

	row := s.pool.QueryRow(ctx, sql, args...)

	var out int64
	if err := row.Scan(&out); err != nil {
		return 0, fmt.Errorf("row.Scan: %w", err)
	}

	return out, nil
}

func (s *Storage) PutDeviceStatus(ctx context.Context, id int64, status entities.Status) error {
	// TODO implement
	return nil
}

func (s *Storage) SetDeviceCommands(ctx context.Context, id int64, commands entities.Commands) error {
	raw, err := json.Marshal(commands)
	if err != nil {
		return fmt.Errorf("json.Marshal commands=%v: %w", commands, err)
	}

	sql, args, err := psql.Insert("device_configuration").Columns(
		"id",
		"data",
	).Values(id, string(raw)).
		Suffix("on conflict (id) do update set data=?", string(raw)).ToSql()
	if err != nil {
		return fmt.Errorf("psql.ToSql: %w", err)
	}

	if _, err := s.pool.Exec(ctx, sql, args...); err != nil {
		return fmt.Errorf("pool.Exec: %w", err)
	}

	return nil
}

func (s *Storage) GetDeviceCommands(ctx context.Context, id int64) (entities.Commands, error) {
	sql, args, err := psql.Select("d.type_id", "dc.data").
		From("device_configuration as dc").
		Join("device as d on dc.id=d.id").
		Where("dc.id = ?", id).ToSql()
	if err != nil {
		return nil, fmt.Errorf("psql.ToSql: %w", err)
	}

	row := s.pool.QueryRow(ctx, sql, args...)

	var (
		deviceTypeID entities.TypeID
		raw          []byte
	)
	if err := row.Scan(
		&deviceTypeID,
		&raw,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, storage.ErrNotFound
		}

		return nil, fmt.Errorf("row.Scan: %w", err)
	}

	switch deviceTypeID {
	case entities.HeatingSystemTypeID:
		var commands entities.TemperatureCommands
		if err := json.Unmarshal(raw, &commands); err != nil {
			return nil, fmt.Errorf("json.Unmarshal raw=%s: %w", string(raw), err)
		}
		return commands, nil
	default:
		return nil, fmt.Errorf("unspecified device_configuration.data raw=%s", string(raw))
	}
}
