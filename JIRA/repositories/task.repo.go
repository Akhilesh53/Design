package repositories

import (
	"errors"
	"pattern/JIRA/entities"
	"sync"
)

var taskRepOnce sync.Once
var taskRepoInstance *TaskRepository

// TaskRepository struct
type TaskRepository struct {
	taskMap map[uint16]*entities.Task
}

// NewTaskRepository function
func NewTaskRepository() *TaskRepository {
	taskRepOnce.Do(func() {
		taskRepoInstance = &TaskRepository{
			taskMap: make(map[uint16]*entities.Task),
		}
	})
	return taskRepoInstance
}

// Save function with input as task entity
func (r *TaskRepository) CreateTask(task *entities.Task) (*entities.Task, error) {
	r.taskMap[task.GetTaskID()] = task
	return task, nil
}

// GetTask function with input as task id
func (r *TaskRepository) GetTask(taskID uint16) (*entities.Task, error) {
	if task, ok := r.taskMap[taskID]; ok {
		return task, nil
	}
	return nil, errors.New("task not found")
}

// FindAll function
func (r *TaskRepository) FindAll() []*entities.Task {
	tasks := make([]*entities.Task, 0)
	for _, task := range r.taskMap {
		tasks = append(tasks, task)
	}
	return tasks
}

// Delete function with input as task id
func (r *TaskRepository) Delete(taskID uint16) {
	delete(r.taskMap, taskID)
}

// UpdateTask function with input as task entity
func (r *TaskRepository) UpdateTask(task *entities.Task) (*entities.Task, error) {
	r.taskMap[task.GetTaskID()] = task
	return task, nil
}
