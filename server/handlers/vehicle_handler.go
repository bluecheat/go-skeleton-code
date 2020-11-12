package handlers

import (
	"context"
	empty "github.com/golang/protobuf/ptypes/empty"
	"skeleton-code/components/vehicle"
	"skeleton-code/proto/generated"
)

var (
	ErrorInvalidArgument = "invalid argument = %s"
)

type vehicleHandler struct {
	vc *vehicle.VehicleComponent
}

func NewVehicleHandler(vc *vehicle.VehicleComponent) *vehicleHandler {
	return &vehicleHandler{
		vc: vc,
	}
}

func (v vehicleHandler) RegisterVehicle(ctx context.Context, request *generated.RegisterVehicleRequest) (*generated.Vehicle, error) {
	v.vc.RegisterVehicle()
	return nil, nil
}

func (v vehicleHandler) ListVehicle(ctx context.Context, request *generated.ListVehicleRequest) (*generated.Vehicles, error) {
	v.vc.ListVehicle()
	return nil, nil
}

func (v vehicleHandler) GetVehicle(ctx context.Context, id *generated.VehicleID) (*generated.Vehicle, error) {
	v.vc.GetVehicle()
	return nil, nil
}

func (v vehicleHandler) UpdateVehicle(ctx context.Context, id *generated.VehicleID) (*generated.Vehicle, error) {
	v.vc.UpdateVehicle()
	return nil, nil
}

func (v vehicleHandler) DeleteVehicle(ctx context.Context, id *generated.VehicleID) (*empty.Empty, error) {
	v.vc.DeleteVehicle()
	return nil, nil
}
