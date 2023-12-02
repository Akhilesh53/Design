package interfaces

type Device interface{
	TogglePower() 
	GetDeviceType() string
	EnablePower() 
	DisablePower()
}