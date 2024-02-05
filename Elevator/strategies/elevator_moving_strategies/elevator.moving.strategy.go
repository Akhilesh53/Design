package elevatormovingstrategies

import "pattern/Elevator/models"

type ElevatorMovingStrategy interface {
	MoveElevator(elevator *models.ElevatorCar)
}
