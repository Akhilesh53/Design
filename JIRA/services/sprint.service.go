package services

import (
	"pattern/JIRA/entities"
	"pattern/JIRA/repositories"
	"sync"
)

var sprintServiceOnce sync.Once
var sprintService *SprintService

// SprintService struct
type SprintService struct {
	sprintRepository *repositories.SprintRepository
	taskReposotiry   *repositories.TaskRepository
}

// NewSprintService function
func NewSprintService() *SprintService {
	sprintServiceOnce.Do(func() {
		sprintService = &SprintService{
			sprintRepository: repositories.NewSprintRepository(),
			taskReposotiry:   repositories.NewTaskRepository(),
		}
	})
	return sprintService
}

// CreateSprint function with input as sprint dto
func (ss *SprintService) CreateSprint(sprintName string, startDate string, endDate string, user entities.User) (*entities.Sprint, error) {
	sprint := entities.NewSprint(sprintName, startDate, endDate, user)
	return ss.sprintRepository.CreateSprint(sprint)
}

// GetSprint function with input as GetSprintDto
func (ss *SprintService) GetSprint(sprintId uint16) (*entities.Sprint, error) {
	return ss.sprintRepository.GetSprint(sprintId)
}

// AddTaskToSprint function with input as AddTaskToSprintDto
func (ss *SprintService) AddTaskToSprint(sprintID uint16, taskID uint16) (*entities.Sprint, error) {
	sprint, err := ss.sprintRepository.GetSprint(sprintID)
	if err != nil {
		return nil, err
	}
	task, err := ss.taskReposotiry.GetTask(taskID)
	if err != nil {
		return nil, err
	}

	sprint.AddTask(*task)
	return ss.sprintRepository.Save(sprint)
}
