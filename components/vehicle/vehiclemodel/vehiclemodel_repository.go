package vehiclemodel

import (
	"context"
	"go.uber.org/fx"
	"skeleton-code/config"
	"skeleton-code/database"
	"time"
)

// Vehicle Model Interface
type IVehicleModelRepository interface {
	registerVehicleModel(ctx context.Context, model *VehicleModel) (*VehicleModel, error)
	getVehicleModel(ctx context.Context, filter *VehicleModel) (*VehicleModel, error)
}

func NewVehicleModelRepository(lifecycle fx.Lifecycle, config *config.Config, db database.Database) IVehicleModelRepository {
	switch config.Database.Driver {
	case database.DB_MYSQL:
		return &vehicleModelMariaDBRepository{db}
	default:
		return &vehicleModelMemRepository{db}
	}
}

type vehicleModelMemRepository struct {
	db database.Database
}

func (v vehicleModelMemRepository) registerVehicleModel(ctx context.Context, model *VehicleModel) (*VehicleModel, error) {
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

func (v vehicleModelMemRepository) getVehicleModel(ctx context.Context, model *VehicleModel) (*VehicleModel, error) {
	val, err := v.db.Get(model, model.ID)
	if err != nil {
		return nil, err
	}
	vehicle := val.(*VehicleModel)
	return vehicle, nil
}

type vehicleModelMariaDBRepository struct {
	db database.Database
}

func (v vehicleModelMariaDBRepository) registerVehicleModel(ctx context.Context, model *VehicleModel) (*VehicleModel, error) {

	if err := v.db.Set(model, "id = ?", model.ID); err != nil {
		return nil, err
	}
	return model, nil
}

func (v vehicleModelMariaDBRepository) getVehicleModel(ctx context.Context, model *VehicleModel) (*VehicleModel, error) {
	val, err := v.db.Get(model, "id = ?", model.ID)
	if err != nil {
		return nil, err
	}
	vehicle := val.(*VehicleModel)
	return vehicle, nil
}
