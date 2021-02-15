package handlers

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/status"
	"skeleton-code/components"
	"skeleton-code/components/vehicle"
	"skeleton-code/errors"
	"skeleton-code/proto/generated"
	"skeleton-code/utils"
)

type vehicleHandler struct {
	vc vehicle.IVehicleService
}

func NewVehicleHandler(ctx components.Context) *vehicleHandler {
	return &vehicleHandler{
		vc: ctx.GetVehicleService(),
	}
}

func (v vehicleHandler) RegisterVehicle(ctx context.Context, request *generated.RegisterVehicleRequest) (*generated.Vehicle, error) {
	if request.Name == "" {
		return nil, errors.Error("required vehicle Name ", errors.ValidationErrCode)
	}
	if request.Vin == "" {
		return nil, errors.Error("required vehicle Vin ", errors.ValidationErrCode)
	}
	if request.Number == "" {
		return nil, errors.Error("required vehicle Number ", errors.ValidationErrCode)
	}

	vehicle, err := v.vc.RegisterVehicle(ctx, request)
	if err != nil {
		return nil, status.Convert(err).Err()
	}

	return &generated.Vehicle{
		Id:        vehicle.ID,
		Name:      "-",
		Vin:       vehicle.VIN,
		Number:    vehicle.PlateNumber,
		Status:    0,
		CreatedAt: utils.TimeToDateString(vehicle.CreatedAt),
		UpdatedAt: utils.TimeToDateString(vehicle.UpdatedAt),
	}, nil
}

func (v vehicleHandler) GetVehicle(ctx context.Context, request *generated.VehicleID) (*generated.Vehicle, error) {
	if request.Id == 0 {
		return nil, errors.Error("required vehicle id ", errors.ValidationErrCode)
	}

	vehicle, err := v.vc.GetVehicle(ctx, request)
	if err != nil {
		return nil, status.Convert(err).Err()
	}

	return &generated.Vehicle{
		Id:        vehicle.ID,
		Name:      "-",
		Vin:       vehicle.VIN,
		Number:    vehicle.PlateNumber,
		Status:    0,
		CreatedAt: utils.TimeToDateString(vehicle.CreatedAt),
		UpdatedAt: utils.TimeToDateString(vehicle.UpdatedAt),
	}, nil
}

func (v vehicleHandler) UpdateVehicle(ctx context.Context, request *generated.UpdateVehicleRequest) (*empty.Empty, error) {
	if request.Id == 0 {
		return nil, errors.Error("required vehicle id ", errors.ValidationErrCode)
	}
	if request.Name == "" {
		return nil, errors.Error("required vehicle Name ", errors.ValidationErrCode)
	}
	if request.Vin == "" {
		return nil, errors.Error("required vehicle Vin ", errors.ValidationErrCode)
	}
	if request.Number == "" {
		return nil, errors.Error("required vehicle Number ", errors.ValidationErrCode)
	}

	_, err := v.vc.UpdateVehicle(ctx, request)
	if err != nil {
		return nil, status.Convert(err).Err()
	}

	return &empty.Empty{}, nil
}

func (v vehicleHandler) DeleteVehicle(ctx context.Context, request *generated.VehicleID) (*empty.Empty, error) {
	if request.Id == 0 {
		return nil, errors.Error("required vehicle id ", errors.ValidationErrCode)
	}

	_, err := v.vc.DeleteVehicle(ctx, request)
	if err != nil {
		return nil, status.Convert(err).Err()
	}
	return &empty.Empty{}, nil
}
