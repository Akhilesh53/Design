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
func (tc *TaskController) CreateTask(createTaskDto *dtos.CreateTaskDto) (*entities.Task, error) {
	return tc.taskService.CreateTask(createTaskDto.GetTitle(), createTaskDto.GetDescription(), createTaskDto.GetCreatedAt(), createTaskDto.GetCreatedBy(), createTaskDto.GetTaskType())
}

// GetTask function with input as GetTaskDto
func (tc *TaskController) GetTask(getTaskDto *dtos.GetTaskDto) (*entities.Task, error) {
	return tc.taskService.GetTask(getTaskDto.GetTaskId())
}

// Add Assignee to task
func (tc *TaskController) AddAssignee(addAssigneeTaskDto *dtos.AddAssigneeTaskDto) (*entities.Task, error) {
	return tc.taskService.AddAssignee(addAssigneeTaskDto.GetTaskId(), addAssigneeTaskDto.GetAssignee())
}

// chnage task status
func (tc *TaskController) ChangeTaskStatus(changeTaskStatusDto *dtos.ChangeTaskStatusDto) (*entities.Task, error) {
	return tc.taskService.ChangeTaskStatus(changeTaskStatusDto.GetTaskID(), changeTaskStatusDto.GetStatus())
}

// GetTasksByAssignee
func (tc *TaskController) GetTasksByAssignee(getTasksByAssigneeDto *dtos.GetTasksByAssigneeDto) ([]entities.Task, error) {
	return tc.taskService.GetTasksByAssignee(getTasksByAssigneeDto.GetAssigneeId())
}