package services

import (
	"pattern/JIRA/dtos"
	"pattern/JIRA/entities"
	"sync"
)

var taskServiceOnce sync.Once
var taskService *TaskService

// TaskService struct
type TaskService struct {
	taskRepository repositories.TaskRepository
}

// NewTaskService function
func NewTaskService() *TaskService {
	taskServiceOnce.Do(func() {
		taskService = &TaskService{
			taskRepository: repositories.NewTaskRepository(),
		}
	})
	return taskService
}

// CreateTask function with input as task dto
func (ts *TaskService) CreateTask(createTaskDto dtos.CreateTaskDto) (*entities.Task, error) {
	return ts.taskRepository.CreateTask(createTaskDto)
}

// GetTask function with input as GetTaskDto
func (ts *TaskService) GetTask(getTaskDto dtos.GetTaskDto) (*entities.Task, error) {
	return ts.taskRepository.GetTask(getTaskDto.GetTaskId())
}
