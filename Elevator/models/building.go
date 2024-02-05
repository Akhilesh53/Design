package models

type Building struct {
	numFloors    int32
	numElevators int32
	floors       []*Floor
	elavators    []*ElevatorCar
}
