package elevatorchossingstrategy

import (
	"pattern/Elevator/models"
)

type ElevatorChoosingStrategy interface {
	ChooseElevator(panel *models.ExternalPanel) uint32
}
