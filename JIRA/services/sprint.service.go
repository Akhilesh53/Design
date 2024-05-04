package services

import (
	"pattern/JIRA/dtos"
	"pattern/JIRA/entities"
	"pattern/JIRA/repositories"
	"sync"
)

var sprintServiceOnce sync.Once
var sprintService *SprintService

// SprintService struct
type SprintService struct {
	sprintRepository *repositories.SprintRepository
}

// NewSprintService function
func NewSprintService() *SprintService {
	sprintServiceOnce.Do(func() {
		sprintService = &SprintService{
			sprintRepository: repositories.NewSprintRepository(),
		}
	})
	return sprintService
}

// CreateSprint function with input as sprint dto
func (ss *SprintService) CreateSprint(createSprintDto dtos.CreateSprintDto) (*entities.Sprint, error) {
	return ss.sprintRepository.CreateSprint(createSprintDto)
}

// GetSprint function with input as GetSprintDto
func (ss *SprintService) GetSprint(getSprintDto dtos.GetSprintDto) (*entities.Sprint, error) {
	return ss.sprintRepository.GetSprint(getSprintDto.GetSprintID())
}
