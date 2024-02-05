package controller

import (
	"pattern/Elevator/models"
	elevatorchossingstrategy "pattern/Elevator/strategies/elevator_choosing_strategy"
)

type ExternalPanelController struct {
	elevatorChoosingStrategy elevatorchossingstrategy.ElevatorChoosingStrategy
	externalPanel            *models.ExternalPanel
}

func (c *ExternalPanelController) GetPanel() *models.ExternalPanel {
	return c.externalPanel
}

func (c *ExternalPanelController) ChooseElevator() uint32 {
	return c.elevatorChoosingStrategy.ChooseElevator(c.externalPanel)
}
