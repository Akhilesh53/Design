package services

import (
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

func (ts *TaskService) CreateTask(title, description string, createdAt time.Time, createdBy entities.User, taskType entities.TaskType) (*entities.Task, error) {
	task := entities.NewTask(title, description, createdAt, createdBy, taskType)
	return ts.taskRepository.CreateTask(task)
}

// GetTask function with input as GetTaskDto
func (ts *TaskService) GetTask(taskId uint16) (*entities.Task, error) {
	return ts.taskRepository.GetTask(taskId)
}

func (ts *TaskService) AddAssignee(taskId uint16, assignee *entities.User) (*entities.Task, error) {
	task, err := ts.taskRepository.GetTask(taskId)
	if err != nil {
		return nil, err
	}
	task.AddAssignee(*assignee)
	return ts.taskRepository.UpdateTask(task)
}

func (ts *TaskService) ChangeTaskStatus(taskId uint16, status entities.TaskStatus) (*entities.Task, error) {
	task, err := ts.taskRepository.GetTask(taskId)
	if err != nil {
		return nil, err
	}
	task.SetStatus(status)
	return ts.taskRepository.UpdateTask(task)
}

// GetTasksByAssigneeDto
func (ts *TaskService) GetTasksByAssignee(assigneeId uint16) ([]entities.Task, error) {
	tasks := ts.taskRepository.FindAll()
	assigneeTasks := make([]entities.Task, 0)
	for _, task := range tasks {
		for _, assignee := range task.GetAssignee() {
			if assignee.GetUserID() == assigneeId {
				assigneeTasks = append(assigneeTasks, *task)
			}
		}
	}
	return assigneeTasks, nil
}
