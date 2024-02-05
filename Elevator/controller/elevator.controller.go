package controller

import (
	"pattern/Elevator/models"
	elevatormovingstrategies "pattern/Elevator/strategies/elevator_moving_strategies"
)

type ElevatorController struct {
	upsideReq              []uint32 // priority queue : smaller floor will come first
	downsideReq            []uint32 // priority queue : smaller floor will come last
	elevatorMovingStrategy elevatormovingstrategies.ElevatorMovingStrategy
	elevator               *models.ElevatorCar
}

func (c *ElevatorController) GetElevator() *models.ElevatorCar {
	return c.elevator
}

func (c *ElevatorController) SubmitInternalRequest(destination uint32) {
	if c.GetElevator().GetCurrFloor() > int(destination) {
		c.downsideReq = append(c.downsideReq, destination)
	}
	if c.GetElevator().GetCurrFloor() <= int(destination) {
		c.upsideReq = append(c.upsideReq, destination)
	}
}

func (c *ElevatorController) Move() {
	c.elevatorMovingStrategy.MoveElevator(c.elevator)
}
