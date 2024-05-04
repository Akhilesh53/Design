package controllers

import (
	"pattern/JIRA/dtos"
	"pattern/JIRA/entities"
	"pattern/JIRA/services"
	"sync"
)

var sprintcontrollerOnce sync.Once
var sprintcontroller *SprintController

// SprintController struct
type SprintController struct {
	sprintService *services.SprintService
}

// NewSprintController function
func NewSprintController() *SprintController {
	sprintcontrollerOnce.Do(func() {
		sprintcontroller = &SprintController{
			sprintService: services.NewSprintService(),
		}
	})
	return sprintcontroller
}

// CreateSprint function with input as sprint dto
func (sc *SprintController) CreateSprint(createSprintDto *dtos.CreateSprintDto) (*entities.Sprint, error) {
	return sc.sprintService.CreateSprint(createSprintDto.GetName(), createSprintDto.GetStartDate(), createSprintDto.GetEndDate(), createSprintDto.GetCreatedBy())
}

// GetSprint function with input as GetSprintDto
func (sc *SprintController) GetSprint(getSprintDto *dtos.GetSprintDto) (*entities.Sprint, error) {
	return sc.sprintService.GetSprint(getSprintDto.GetSprintID())
}

// / add task
func (sc *SprintController) AddTaskToSprint(addTaskToSprintDto *dtos.AddTaskToSprintDto) (*entities.Sprint, error) {
	return sc.sprintService.AddTaskToSprint(addTaskToSprintDto.GetSprintID(), addTaskToSprintDto.GetTaskID())
}
