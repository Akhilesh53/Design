package entities

import "time"

var sprintId uint16

func init() {
	sprintId = 0
}

func getSprintId() uint16 {
	sprintId++
	return sprintId
}

type Sprint struct {
	sprintid  uint16
	name      string
	startdate string
	enddate   string
	tasks     []Task
	createdAt time.Time
	createdBy User
}

func NewSprint(name string, startdate string, enddate string, createdBy User) *Sprint {
	return &Sprint{
		sprintid:  getSprintId(),
		name:      name,
		startdate: startdate,
		enddate:   enddate,
		createdAt: time.Now(),
		createdBy: createdBy,
	}
}

func NewDefaultSprint(user User) *Sprint {
	return &Sprint{
		sprintid:  getSprintId(),
		createdAt: time.Now(),
		createdBy: user,
	}
}

func (s *Sprint) GetSprintID() uint16 {
	return s.sprintid
}

func (s *Sprint) GetName() string {
	return s.name
}

func (s *Sprint) GetStartDate() string {
	return s.startdate
}

func (s *Sprint) GetEndDate() string {
	return s.enddate
}

func (s *Sprint) GetTasks() []Task {
	return s.tasks
}

func (s *Sprint) SetSprintID(sprintid uint16) *Sprint {
	s.sprintid = sprintid
	return s
}

func (s *Sprint) SetName(name string) *Sprint {
	s.name = name
	return s
}

func (s *Sprint) SetStartDate(startdate string) *Sprint {
	s.startdate = startdate
	return s
}

func (s *Sprint) SetEndDate(enddate string) *Sprint {
	s.enddate = enddate
	return s
}

func (s *Sprint) SetTasks(tasks []Task) *Sprint {
	s.tasks = tasks
	return s
}

func (s *Sprint) AddTask(task Task) *Sprint {
	s.tasks = append(s.tasks, task)
	return s
}

func (s *Sprint) RemoveTask(task Task) *Sprint {
	for i, t := range s.tasks {
		if t.GetTaskID() == task.GetTaskID() {
			s.tasks = append(s.tasks[:i], s.tasks[i+1:]...)
			break
		}
	}
	return s
}
