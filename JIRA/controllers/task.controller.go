package controllers

import (
	"pattern/JIRA/dtos"
	"pattern/JIRA/entities"
	"pattern/JIRA/services"
	"sync"
)

var taskcontollerOnce sync.Once
var taskcontroller *TaskController

// TaskController struct
type TaskController struct {
	taskService *services.TaskService
}

// NewTaskController function
func NewTaskController() *TaskController {
	taskcontollerOnce.Do(func() {
		taskcontroller = &TaskController{
			taskService: services.NewTaskService(),
		}
	})
	return taskcontroller
}

// CreateTask function with input as task dto
func (tc *TaskController) CreateTask(createTaskDto dtos.CreateTaskDto) (*entities.Task, error) {
	return tc.taskService.CreateTask(createTaskDto)
}

// GetTask function with input as GetTaskDto
func (tc *TaskController) GetTask(getTaskDto dtos.GetTaskDto) (*entities.Task, error) {
	return tc.taskService.GetTask(getTaskDto)
}
