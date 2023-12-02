package entities

type acDevice struct {
	deviceType string
	power      bool
}

func NewACDevice() *acDevice {
	return &acDevice{
		deviceType: "AC",
		power:      false,
	}
}

func (d *acDevice) GetDeviceType() string {
	return d.deviceType
}

func (d *acDevice) IsEnabled() bool {
	return d.power == true
}

func (d *acDevice) EnablePower() {
	d.power = true
}

func (d *acDevice) DisablePower() {
	d.power = false
}

func (d *acDevice) TogglePower() {
	if d.IsEnabled() {
		d.DisablePower()
		return
	}
	d.DisablePower()
}
