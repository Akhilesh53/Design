package entities

import "time"

var taskId uint16

func init() {
	taskId = 0
}

func getTaskId() uint16 {
	taskId++
	return taskId
}

type Task struct {
	taskid      uint16
	title       string
	description string
	assignee    []User
	createdBy   User
	createdAt   time.Time
	taskType    TaskType
	status      TaskStatus
	dueDate     time.Time
}

func NewTask(title string, description string, createdAt time.Time, createdBy User, taskType TaskType) *Task {
	return &Task{
		taskid:      getTaskId(),
		title:       title,
		description: description,
		assignee:    make([]User, 0),
		createdAt:   createdAt,
		createdBy:   createdBy,
		taskType:    taskType,
		status:      OPEN,
	}
}

func NewDefaultTask() *Task {
	return &Task{
		taskid:    getTaskId(),
		status:    OPEN,
		createdAt: time.Now(),
	}
}

func (t *Task) GetTaskID() uint16 {
	return t.taskid
}

func (t *Task) GetTitle() string {
	return t.title
}

func (t *Task) GetDescription() string {
	return t.description
}

func (t *Task) GetAssignee() []User {
	return t.assignee
}

func (t *Task) GetCreatedAt() time.Time {
	return t.createdAt
}

func (t *Task) GetTaskType() TaskType {
	return t.taskType
}

func (t *Task) GetStatus() TaskStatus {
	return t.status
}

func (t *Task) GetCreatedBy() User {
	return t.createdBy
}

func (t *Task) SetTaskID(taskid uint16) *Task {
	t.taskid = taskid
	return t
}

func (t *Task) SetTitle(title string) *Task {
	t.title = title
	return t
}

func (t *Task) SetDescription(description string) *Task {
	t.description = description
	return t
}

func (t *Task) SetAssignee(assignees []User) *Task {
	t.assignee = assignees
	return t
}

func (t *Task) SetCreatedBy(createdBy User) *Task {
	t.createdBy = createdBy
	return t
}

func (t *Task) AddAssignee(assignee User) *Task {
	t.assignee = append(t.assignee, assignee)
	return t
}

func (t *Task) SetCreatedAt(createdAt time.Time) *Task {
	t.createdAt = createdAt
	return t
}

func (t *Task) SetTaskType(taskType TaskType) *Task {
	t.taskType = taskType
	return t
}

func (t *Task) SetStatus(status TaskStatus) *Task {
	t.status = status
	return t
}

func (t *Task) CreateTask() *Task {
	t.status = OPEN
	t.createdAt = time.Now()
	return t
}
