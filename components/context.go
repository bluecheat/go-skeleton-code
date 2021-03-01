package components

import (
	"skeleton-code/components/member"
	"skeleton-code/components/vehicle"
	"skeleton-code/components/vehicle/vehiclemodel"
)

type Context interface {
	GetVehicleService() vehicle.IVehicleService
	GetVehicleModelService() vehiclemodel.IVehicleService
	GetMemberService() member.IMemberService
}

type ComponentContext struct {
	vs  vehicle.IVehicleService
	vsm vehiclemodel.IVehicleService
	ms  member.IMemberService
}

func NewComponentContext(vs vehicle.IVehicleService, vsm vehiclemodel.IVehicleService, ms member.IMemberService) Context {
	return &ComponentContext{
		vs:  vs,
		vsm: vsm,
		ms:  ms,
	}
}

func (cc *ComponentContext) GetVehicleService() vehicle.IVehicleService {
	return cc.vs
}

func (cc *ComponentContext) GetVehicleModelService() vehiclemodel.IVehicleService {
	return cc.vsm
}

func (cc *ComponentContext) GetMemberService() member.IMemberService {
	return cc.ms
}
