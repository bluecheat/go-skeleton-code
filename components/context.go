package components

import "skeleton-code/components/vehicle"

type Context interface {
	GetVehicleComponent() *vehicle.VehicleComponent
}

type components struct {
	vehicleComponent *vehicle.VehicleComponent
}

func NewContext(vehicleComponent *vehicle.VehicleComponent) Context {
	return &components{
		vehicleComponent: vehicleComponent,
	}
}

func (c *components) GetVehicleComponent() *vehicle.VehicleComponent {
	return c.vehicleComponent
}
