package services

import (
	"pattern/JIRA/dtos"
	"pattern/JIRA/entities"
	"pattern/JIRA/repositories"
	"sync"
	"time"
)

var taskServiceOnce sync.Once
var taskService *TaskService

// TaskService struct
type TaskService struct {
	taskRepository *repositories.TaskRepository
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
func (ts *TaskService) CreateTask(title string, description string, createdAt time.Time, createdBy entities.User, taskType entities.TaskType) (*entities.Task, error) {
	task := entities.NewTask(title, description, createdAt, createdBy, taskType)
	return ts.taskRepository.CreateTask(task)
}

// GetTask function with input as GetTaskDto
func (ts *TaskService) GetTask(getTaskDto dtos.GetTaskDto) (*entities.Task, error) {
	return ts.taskRepository.GetTask(getTaskDto.GetTaskId())
}
