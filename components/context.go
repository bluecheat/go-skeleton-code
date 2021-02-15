package components

import (
	"skeleton-code/components/member"
	"skeleton-code/components/vehicle"
)

type Context interface {
	GetVehicleService() vehicle.IVehicleService
	GetMemberService() member.IMemberService
}

type ComponentContext struct {
	vs vehicle.IVehicleService
	ms member.IMemberService
}

func NewComponentContext(vs vehicle.IVehicleService, ms member.IMemberService) Context {
	return &ComponentContext{
		vs: vs,
		ms: ms,
	}
}

func (cc *ComponentContext) GetVehicleService() vehicle.IVehicleService {
	return cc.vs
}

func (cc *ComponentContext) GetMemberService() member.IMemberService {
	return cc.ms
}
