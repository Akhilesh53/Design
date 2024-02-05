package elevatorchossingstrategy

import "pattern/Elevator/models"

type RandomElevatorChoosingStrategy struct{}

func (r *RandomElevatorChoosingStrategy) ChooseElevator(panel *models.ExternalPanel) {
	// a control panel has access to two different elevators
	// choose one elevator based on the strategy
	// might return elevator id
}
