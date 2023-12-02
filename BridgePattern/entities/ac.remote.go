package entities

import "pattern/BridgePattern/interfaces"

type acRemote struct {
	device      interfaces.Device
	power       bool
}

func NewACRemote(device interfaces.Device) *acRemote {
	return &acRemote{
		device:      device,
		power:       false,
	}
}

func (r *acRemote) TogglePower() {
	if r.power {
		r.power = false
	} else {
		r.power = true
	}
}

func (r *acRemote) IsPowerEnabled() bool {
	return r.power == true
}

func (r *acRemote) EnablePower() {
	r.power = true
}

func (r *acRemote) DisablePower() {
	r.power = false
}
