package vehicle

import (
	"context"
	"go.uber.org/fx"
	"skeleton-code/config"
	"skeleton-code/database"
	"time"
)

// Vehicle Interface
type IVehicleRepository interface {
	registerVehicle(ctx context.Context, model *Vehicle) (*Vehicle, error)
	getVehicle(ctx context.Context, filter *Vehicle) (*Vehicle, error)
	updateVehicle(ctx context.Context, model *Vehicle) (bool, error)
	deleteVehicle(ctx context.Context, filter *Vehicle) (bool, error)
}

func NewVehicleRepository(lifecycle fx.Lifecycle, config *config.Config, db database.Database) IVehicleRepository {
	switch config.Database.Driver {
	case "mysql":
		return &vehicleMariaDBRepository{db}
	default:
		return &vehicleMemRepository{db}
	}
}

type vehicleMemRepository struct {
	db database.Database
}

func (v vehicleMemRepository) registerVehicle(ctx context.Context, model *Vehicle) (*Vehicle, error) {
	cnt, err := v.db.Count(model)
	if err != nil {
		return nil, err
	}
	model.ID = uint64(cnt) + 1
	model.CreatedAt = time.Now()
	model.UpdatedAt = time.Now()
	if err := v.db.Set(model, model.ID); err != nil {
		return nil, err
	}
	return model, nil
}

func (v vehicleMemRepository) getVehicle(ctx context.Context, model *Vehicle) (*Vehicle, error) {
	val, err := v.db.Get(model, model.ID)
	if err != nil {
		return nil, err
	}
	vehicle := val.(*Vehicle)
	return vehicle, nil
}

func (v vehicleMemRepository) updateVehicle(ctx context.Context, model *Vehicle) (bool, error) {
	val, err := v.db.Get(model, model.ID)
	if err != nil {
		return false, err
	}
	vehicle := val.(*Vehicle)

	model.CreatedAt = vehicle.CreatedAt
	model.UpdatedAt = time.Now()

	if err := v.db.Update(model, model.ID); err != nil {
		return false, err
	}
	return true, nil
}

func (v vehicleMemRepository) deleteVehicle(ctx context.Context, model *Vehicle) (bool, error) {
	if err := v.db.Delete(model, model.ID); err != nil {
		return false, err
	}
	return true, nil
}

type vehicleMariaDBRepository struct {
	db database.Database
}

func (v vehicleMariaDBRepository) registerVehicle(ctx context.Context, model *Vehicle) (*Vehicle, error) {

	if err := v.db.Set(model, "id = ?", model.ID); err != nil {
		return nil, err
	}
	return model, nil
}

func (v vehicleMariaDBRepository) getVehicle(ctx context.Context, model *Vehicle) (*Vehicle, error) {
	val, err := v.db.Get(model, "id = ?", model.ID)
	if err != nil {
		return nil, err
	}
	vehicle := val.(*Vehicle)
	return vehicle, nil
}

func (v vehicleMariaDBRepository) updateVehicle(ctx context.Context, model *Vehicle) (bool, error) {
	val, err := v.db.Get(model, "id = ?", model.ID)
	if err != nil {
		return false, err
	}
	vehicle := val.(*Vehicle)

	model.CreatedAt = vehicle.CreatedAt
	model.UpdatedAt = time.Now()

	if err := v.db.Update(model, "id = ?", model.ID); err != nil {
		return false, err
	}
	return true, nil
}

func (v vehicleMariaDBRepository) deleteVehicle(ctx context.Context, model *Vehicle) (bool, error) {
	if err := v.db.Delete(model, "id = ?", model.ID); err != nil {
		return false, err
	}
	return true, nil
}
