package models

import "pattern/Elevator/enums"

type ElevatorCar struct {
	elevatorId      uint32
	currFloor       int
	movingDirection enums.Direction
	elevatorStatus  enums.Status
	door            *Door
	buttons         []*InternalButton
	maxFloor        int
	minFloor        int
}

func (elevator *ElevatorCar) GetElevatorId() uint32 {
	return elevator.elevatorId
}

func (elevator *ElevatorCar) GetMovingDirection() enums.Direction {
	return elevator.movingDirection
}

func (elevator *ElevatorCar) GetCurrFloor() int {
	return elevator.currFloor
}