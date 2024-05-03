package dtos

import (
	"pattern/JIRA/entities"
	"time"
)

// Create Task DTO struct
type CreateTaskDto struct {
	title       string
	description string
	createdAt   time.Time
	createdBy   entities.User
	taskType    entities.TaskType
}

// Get Task DTO struct
type GetTaskDto struct {
	taskId uint16
}

// NewCreateTaskDto function
func NewCreateTaskDto(title string, description string, createdAt time.Time, createdBy entities.User, taskType entities.TaskType) *CreateTaskDto {
	return &CreateTaskDto{
		title:       title,
		description: description,
		createdAt:   createdAt,
		createdBy:   createdBy,
		taskType:    taskType,
	}
}

// NewGetTaskDto function
func NewGetTaskDto(taskId uint16) *GetTaskDto {
	return &GetTaskDto{
		taskId: taskId,
	}
}

// GetTitle function
func (ctd *CreateTaskDto) GetTitle() string {
	return ctd.title
}

// GetDescription function
func (ctd *CreateTaskDto) GetDescription() string {
	return ctd.description
}

// GetCreatedAt function
func (ctd *CreateTaskDto) GetCreatedAt() time.Time {
	return ctd.createdAt
}

// GetCreatedBy function
func (ctd *CreateTaskDto) GetCreatedBy() entities.User {
	return ctd.createdBy
}

// GetTaskType function
func (ctd *CreateTaskDto) GetTaskType() entities.TaskType {
	return ctd.taskType
}

// GetTaskId function
func (gtd *GetTaskDto) GetTaskId() uint16 {
	return gtd.taskId
}

// SetTitle function
func (ctd *CreateTaskDto) SetTitle(title string) *CreateTaskDto {
	ctd.title = title
	return ctd
}

// SetDescription function
func (ctd *CreateTaskDto) SetDescription(description string) *CreateTaskDto {
	ctd.description = description
	return ctd
}

// SetCreatedAt function
func (ctd *CreateTaskDto) SetCreatedAt(createdAt time.Time) *CreateTaskDto {
	ctd.createdAt = createdAt
	return ctd
}

// SetCreatedBy function
func (ctd *CreateTaskDto) SetCreatedBy(createdBy entities.User) *CreateTaskDto {
	ctd.createdBy = createdBy
	return ctd
}

// SetTaskType function
func (ctd *CreateTaskDto) SetTaskType(taskType entities.TaskType) *CreateTaskDto {
	ctd.taskType = taskType
	return ctd
}

// SetTaskId function
func (gtd *GetTaskDto) SetTaskId(taskId uint16) *GetTaskDto {
	gtd.taskId = taskId
	return gtd
}
