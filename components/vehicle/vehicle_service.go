package vehicle

import (
	"context"
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
	repository IVehicleRepository
}

func NewVehicleService(repository IVehicleRepository) IVehicleService {
	return &vehicleService{
		repository: repository,
	}
}

// RegisterVehicle 차량 생성
func (v vehicleService) RegisterVehicle(ctx context.Context, request *generated.RegisterVehicleRequest) (*Vehicle, error) {

	newVehicle := &Vehicle{
		Name:        request.Name,
		PlateNumber: request.Number,
		VIN:         request.Vin,
	}
	createdVehicle, err := v.repository.registerVehicle(ctx, newVehicle)
	if err != nil {
		return nil, errors.Error(err.Error(), errors.DatabaseErrCode)
	}

	return createdVehicle, nil
}

// GetVehicle 차량 상세
func (v vehicleService) GetVehicle(ctx context.Context, request *generated.VehicleID) (*Vehicle, error) {

	findVehicle, err := v.repository.getVehicle(ctx, &Vehicle{
		ID: request.Id,
	})
	if err != nil {
		return nil, errors.Error("not found vehicle", errors.DatabaseErrCode)
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
	_, err := v.repository.updateVehicle(ctx, updatedVehicle)
	if err != nil {
		return false, errors.Error(err.Error(), errors.DatabaseErrCode)
	}
	return true, nil
}

// DeleteVehicle 차량 삭제
func (v vehicleService) DeleteVehicle(ctx context.Context, request *generated.VehicleID) (bool, error) {
	_, err := v.repository.deleteVehicle(ctx, &Vehicle{
		ID: request.Id,
	})
	if err != nil {
		return false, errors.Error(err.Error(), errors.DatabaseErrCode)
	}
	return true, nil
}
