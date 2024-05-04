package repositories

import (
	"errors"
	"pattern/JIRA/entities"
	"sync"
)

var sprintRepoOnce sync.Once
var sprintRepoInstance *SprintRepository

// SprintRepository struct
type SprintRepository struct {
	sprintMap map[uint16]*entities.Sprint
}

// NewSprintRepository function
func NewSprintRepository() *SprintRepository {
	sprintRepoOnce.Do(func() {
		sprintRepoInstance = &SprintRepository{
			sprintMap: make(map[uint16]*entities.Sprint),
		}
	})
	return sprintRepoInstance
}

// CreateSprint function with input as sprint dto
func (r *SprintRepository) CreateSprint(sprint *entities.Sprint) (*entities.Sprint, error) {
	return r.Save(sprint)
}

// GetSprint function with input as GetSprintDto
func (r *SprintRepository) GetSprint(sprintID uint16) (*entities.Sprint, error) {
	if sprint, ok := r.sprintMap[sprintID]; ok {
		return sprint, nil
	}
	return nil, errors.New("sprint not found")
}

// Save function
func (r *SprintRepository) Save(sprint *entities.Sprint) (*entities.Sprint, error) {
	r.sprintMap[sprint.GetSprintID()] = sprint
	return sprint, nil
}

// FindAll function
func (r *SprintRepository) FindAll() []*entities.Sprint {
	sprints := make([]*entities.Sprint, 0)
	for _, sprint := range r.sprintMap {
		sprints = append(sprints, sprint)
	}
	return sprints
}

// Delete function
func (r *SprintRepository) Delete(sprintID uint16) {
	delete(r.sprintMap, sprintID)
}
