package elevatormovingstrategies

import "pattern/Elevator/models"

type RandomElevatorMovingStrategy struct{}

func (r *RandomElevatorMovingStrategy) MoveElevator(elevator *models.ElevatorCar) {
	// if elevators current direction is upside, move to next upside floor if not at the max floor
	// else reverse the direction

	// vice versa for downside direction
}
