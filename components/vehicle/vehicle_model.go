package vehicle

import "time"

type Vehicle struct {
	ID          uint64
	Name        string
	PlateNumber string
	VIN         string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}
