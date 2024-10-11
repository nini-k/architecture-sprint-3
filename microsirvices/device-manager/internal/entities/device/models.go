package device

import "time"

type TypeID int32

const (
	HeatingSystemTypeID TypeID = 1
)

type Device struct {
	ID           int64
	Name         string
	HouseID      int64
	TypeID       TypeID
	SerialNumber string
	Status       Status
	CreatedAt    time.Time
}

type Status struct {
	IsOn bool
}

type Commands interface{}

type TemperatureCommands struct {
	TargetTemperature int32
}
