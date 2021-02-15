package vehicle

import (
	"context"
	"skeleton-code/database"
	"skeleton-code/errors"
	"skeleton-code/proto/generated"
)

// Vehicle Interface
type IVehicleService interface {
	RegisterVehicle(ctx context.Context, request *generated.RegisterVehicleRequest) (*Vehicle, error)
	GetVehicle(ctx context.Context, id *generated.VehicleID) (*Vehicle, error)
	UpdateVehicle(ctx context.Context, request *generated.UpdateVehicleRequest) (bool, error)
	DeleteVehicle(ctx context.Context, request *generated.VehicleID) (bool, error)
}

type vehicleService struct {
	db database.Database
}

func NewVehicleService(db database.Database) IVehicleService {
	return &vehicleService{
		db: db,
	}
}

// RegisterVehicle 차량 생성
func (v vehicleService) RegisterVehicle(ctx context.Context, request *generated.RegisterVehicleRequest) (*Vehicle, error) {

	newVehicle := &Vehicle{
		Name:        request.Name,
		PlateNumber: request.Number,
		VIN:         request.Vin,
	}

	err := v.db.Set(newVehicle)
	if err != nil {
		return nil, errors.Error(err.Error(), errors.DatabaseErrCode)
	}

	return newVehicle, nil
}

// GetVehicle 차량 상세
func (v vehicleService) GetVehicle(ctx context.Context, id *generated.VehicleID) (*Vehicle, error) {
	findVehicle := &Vehicle{
		ID: id.Id,
	}
	err := v.db.Get(findVehicle)
	if err != nil {
		return nil, errors.Error(err.Error(), errors.DatabaseErrCode)
	}

	return findVehicle, nil
}

// UpdateVehicle 차량 업데이트
func (v vehicleService) UpdateVehicle(ctx context.Context, request *generated.UpdateVehicleRequest) (bool, error) {
	updatedVehicle := &Vehicle{
		ID:          request.Id,
		PlateNumber: request.Number,
		Name:        request.Name,
		VIN:         request.Vin,
	}
	err := v.db.Update(updatedVehicle)
	if err != nil {
		return false, errors.Error(err.Error(), errors.DatabaseErrCode)
	}
	return true, nil
}

// DeleteVehicle 차량 삭제
func (v vehicleService) DeleteVehicle(ctx context.Context, request *generated.VehicleID) (bool, error) {
	err := v.db.Delete(request)
	if err != nil {
		return false, errors.Error(err.Error(), errors.DatabaseErrCode)
	}
	return true, nil
}
