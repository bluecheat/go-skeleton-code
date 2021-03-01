package handlers

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/status"
	"skeleton-code/components"
	"skeleton-code/components/vehicle"
	"skeleton-code/components/vehicle/vehiclemodel"
	"skeleton-code/errors"
	"skeleton-code/logger"
	"skeleton-code/proto/generated"
	"skeleton-code/utils"
)

type vehicleHandler struct {
	vc  vehicle.IVehicleService
	vcm vehiclemodel.IVehicleService
}

func NewVehicleHandler(ctx components.Context) *vehicleHandler {
	return &vehicleHandler{
		vc:  ctx.GetVehicleService(),
		vcm: ctx.GetVehicleModelService(),
	}
}

func (v vehicleHandler) RegisterVehicle(ctx context.Context, request *generated.RegisterVehicleRequest) (*generated.Vehicle, error) {
	if request.Name == "" {
		return nil, errors.Error("required vehicle name ", errors.ValidationErrCode)
	}
	if request.Vin == "" {
		return nil, errors.Error("required vehicle vin ", errors.ValidationErrCode)
	}
	if request.Number == "" {
		return nil, errors.Error("required vehicle number ", errors.ValidationErrCode)
	}

	vehicle, err := v.vc.RegisterVehicle(ctx, request)
	if err != nil {
		return nil, errors.Convert(err)
	}

	logger.Infof("%+v", vehicle)

	return &generated.Vehicle{
		Id:        vehicle.ID,
		Name:      vehicle.Name,
		Vin:       vehicle.VIN,
		Number:    vehicle.PlateNumber,
		Status:    0,
		CreatedAt: utils.TimeToDateString(vehicle.CreatedAt),
	}, nil
}

func (v vehicleHandler) GetVehicle(ctx context.Context, request *generated.VehicleID) (*generated.Vehicle, error) {
	if request.Id == 0 {
		return nil, errors.Error("require vehicle id", errors.ValidationErrCode)
	}

	vehicle, err := v.vc.GetVehicle(ctx, request)
	if err != nil {
		logger.Infof("%+v", errors.Convert(err))
		return nil, errors.Convert(err)
	}

	return &generated.Vehicle{
		Id:        vehicle.ID,
		Name:      vehicle.Name,
		Vin:       vehicle.VIN,
		Number:    vehicle.PlateNumber,
		Status:    0,
		CreatedAt: utils.TimeToDateString(vehicle.CreatedAt),
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

func (v vehicleHandler) RegisterVehicleModel(ctx context.Context, request *generated.RegisterVehicleModelRequest) (*generated.VehicleModel, error) {
	if request.Name == "" {
		return nil, errors.Error("required vehicle model name ", errors.ValidationErrCode)
	}
	if request.Gear == "" {
		return nil, errors.Error("required vehicle model gear ", errors.ValidationErrCode)
	}
	if request.Color == "" {
		return nil, errors.Error("required vehicle model color ", errors.ValidationErrCode)
	}

	vehiclemodel, err := v.vcm.RegisterVehicle(ctx, request)
	if err != nil {
		return nil, errors.Convert(err)
	}

	logger.Infof("%+v", vehiclemodel)

	return &generated.VehicleModel{
		Id:        vehiclemodel.ID,
		Name:      vehiclemodel.Name,
		Gear:      vehiclemodel.Gear,
		Color:     vehiclemodel.Color,
		CreatedAt: utils.TimeToDateString(vehiclemodel.CreatedAt),
	}, nil
}

func (v vehicleHandler) GetVehicleModel(ctx context.Context, request *generated.VehicleModelID) (*generated.VehicleModel, error) {
	if request.Id == 0 {
		return nil, errors.Error("require vehicle Model id", errors.ValidationErrCode)
	}

	vehiclemodel, err := v.vcm.GetVehicle(ctx, request)
	if err != nil {
		logger.Infof("%+v", errors.Convert(err))
		return nil, errors.Convert(err)
	}

	return &generated.VehicleModel{
		Id:        vehiclemodel.ID,
		Name:      vehiclemodel.Name,
		Gear:      vehiclemodel.Gear,
		Color:     vehiclemodel.Color,
		CreatedAt: utils.TimeToDateString(vehiclemodel.CreatedAt),
	}, nil
}
