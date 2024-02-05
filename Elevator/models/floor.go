package models

import "pattern/Elevator/enums"

type Floor struct {
	floorNumber  uint32
	controlPanel *ExternalPanel
}

type ExternalPanel struct {
	panelId             uint32
	direction           enums.Direction
	accessibleElevators []*ElevatorCar // atmost 2
}

func (e *ExternalPanel) GetId() uint32 { return e.panelId }

func (e *ExternalPanel) GetRequestedPanelDirection() enums.Direction { return e.direction }
