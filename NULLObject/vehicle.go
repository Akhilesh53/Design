package main

type IVehicle interface {
	getSeatCount() int
	getFuelCapacity() float64
}

func GetVehicle(vehicleType string) IVehicle {
	if vehicleType == "car" {
		return getCar()
	} else if vehicleType == "bike" {
		return getBike()
	}
	return getDefaultVehicle()
}

// define a car struct and define methods as per interface
type car struct {
	seatCount    int
	fuelCapacity float64
}

func (c *car) getSeatCount() int {
	return c.seatCount
}

func (c *car) getFuelCapacity() float64 {
	return c.fuelCapacity
}

func getCar() IVehicle {
	return &car{
		seatCount:    5,
		fuelCapacity: 25,
	}
}

// define a bike object and define same methiods as per interface
type bike struct {
	seatCount    int
	fuelCapacity float64
}

func (b *bike) getSeatCount() int {
	return b.seatCount
}

func (b *bike) getFuelCapacity() float64 {
	return b.fuelCapacity
}

func getBike() IVehicle {
	return &bike{
		seatCount:    2,
		fuelCapacity: 15,
	}
}

// define a default object and define same methiods as per interface
type defaultVehicle struct {
	seatCount    int
	fuelCapacity float64
}

func (d *defaultVehicle) getSeatCount() int {
	return d.seatCount
}

func (d *defaultVehicle) getFuelCapacity() float64 {
	return d.fuelCapacity
}

func getDefaultVehicle() IVehicle {
	return &defaultVehicle{
		seatCount:    0,
		fuelCapacity: 0,
	}
}
