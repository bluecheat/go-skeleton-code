package vehiclemodel

import (
	"context"
	"skeleton-code/errors"
	"skeleton-code/proto/generated"
)

// Vehicle Interface
type IVehicleService interface {
	RegisterVehicle(ctx context.Context, request *generated.RegisterVehicleModelRequest) (*VehicleModel, error)
	GetVehicle(ctx context.Context, id *generated.VehicleModelID) (*VehicleModel, error)
}

type vehicleService struct {
	repo IVehicleModelRepository
}

func NewVehicleService(repo IVehicleModelRepository) IVehicleService {
	return &vehicleService{
		repo: repo,
	}
}

// RegisterVehicle 차량 생성
func (v vehicleService) RegisterVehicle(ctx context.Context, request *generated.RegisterVehicleModelRequest) (*VehicleModel, error) {

	newVehicle := &VehicleModel{
		Name:  request.Name,
		Color: request.Color,
		Gear:  request.Gear,
	}
	createdVehicle, err := v.repo.registerVehicleModel(ctx, newVehicle)
	if err != nil {
		return nil, errors.Error(err.Error(), errors.DatabaseErrCode)
	}

	return createdVehicle, nil
}

// GetVehicle 차량 상세
func (v vehicleService) GetVehicle(ctx context.Context, request *generated.VehicleModelID) (*VehicleModel, error) {

	findVehicle, err := v.repo.getVehicleModel(ctx, &VehicleModel{
		ID: request.Id,
	})
	if err != nil {
		return nil, errors.Error("not found vehicleModel", errors.DatabaseErrCode)
	}

	return findVehicle, nil
}
