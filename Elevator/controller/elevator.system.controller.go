package controller

import "pattern/Elevator/enums"

type ElevatorSystem struct {
	// list of elevator controllers
	elevatorControllers []*ElevatorController
	// list of external panel controllers
	externalPanelControllers []*ExternalPanelController
	//
}

// when user presses button on external panel, a request will be generated by elevator system
// that this aprticular control panel has been pressed for that direction from this particular floor.
// noew, what external panel controller will do, it will select an elevator based on some strategy assigned
// and assign the request to that particlar elevator. Then elevator will move as per the strategy it has been assigned.
func (elevator *ElevatorSystem) SubmitExternalRequest(direction enums.Direction, floor, panelId uint32) {
	for _, panelController := range elevator.externalPanelControllers {
		if panelController.GetPanel().GetId() == panelId {
			elevatorId := panelController.ChooseElevator()

			//get elevator controller from elevator id
			var elevatorController *ElevatorController
			for _, elevatorController = range elevator.elevatorControllers {
				if elevatorController.GetElevator().GetElevatorId() == elevatorId {
					elevatorController.SubmitInternalRequest(floor)
					break
				}
			}
		}
	}
}

// when a user presses the button from inside lift, elevator controller will get the particular elevator and assign the
// request to that particular elevator
func (elevator *ElevatorSystem) SubmitInternalRequest(destination, elevatorId uint32) {
	for _, elevatorController := range elevator.elevatorControllers {
		if elevatorController.GetElevator().GetElevatorId() == elevatorId {
			elevatorController.SubmitInternalRequest(destination)
		}
	}
}

func (elevator *ElevatorSystem) MoveElevator(elevatorId uint32) {
	for _, elevatorController := range elevator.elevatorControllers {
		if elevatorController.GetElevator().GetElevatorId() == elevatorId {
			elevatorController.Move()
		}
	}
}