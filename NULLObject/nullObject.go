package main

import "fmt"

func main() {
	vehicle := GetVehicle("car")

	fmt.Println(vehicle.getFuelCapacity())
	fmt.Println(vehicle.getSeatCount())

	vehicle = GetVehicle("truck")

	fmt.Println(vehicle.getFuelCapacity())
	fmt.Println(vehicle.getSeatCount())
}
