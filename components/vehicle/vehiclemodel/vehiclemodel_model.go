package vehiclemodel

import "time"

type VehicleModel struct {
	ID        uint64
	Color     string
	Name      string
	Gear      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
