package vehicle

import "time"

type Vehicle struct {
	ID          uint64
	Name        string
	plateNumber string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}
