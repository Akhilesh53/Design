package main

import (
	"fmt"
	"pattern/JIRA/controllers"
	"pattern/JIRA/dtos"
	"pattern/JIRA/entities"
	"time"
)

func main() {

	// declare controllers
	// create a task controller
	taskController := controllers.NewTaskController()
	// create a user controller
	userController := controllers.NewUserController()
	// create a sprint controller
	sprintController := controllers.NewSprintController()

	//=========================================================================================================
	// create users
	createdUser, err := userController.CreateUser(dtos.NewCreateUserDto("Akhilesh Mahajan", "akhileshmahajan@gmail.com"))
	if err != nil {
		fmt.Println("error while creating user : ", err)
		return
	}

	assigneeUser, err := userController.CreateUser(dtos.NewCreateUserDto("Rahul Mahajan", "rahulmahajan@gmail.com"))
	if err != nil {
		fmt.Println("error while creating user : ", err)
		return
	}

	//=========================================================================================================
	// get user
	getUserDto := dtos.NewGetUserDto(createdUser.GetUserID())
	user, err := userController.GetUser(getUserDto)
	if err != nil {
		fmt.Println("error while getting user : ", err)
		return
	}
	fmt.Println("getting user details : ", user.GetUserID(), user.GetName(), user.GetEmail())

	//=========================================================================================================
	// create a task
	task, err := taskController.CreateTask(dtos.NewCreateTaskDto("Task 1", "Task 1 description", time.Now(), *user, entities.FEATURE))

	if err != nil {
		fmt.Println("error while creating task : ", err)
		return
	}

	taskController.AddAssignee(dtos.NewAddAssigneeTaskDto(task.GetTaskID(), assigneeUser))

	//=========================================================================================================

	// get task
	getTaskDto := dtos.NewGetTaskDto(task.GetTaskID())
	task, err = taskController.GetTask(getTaskDto)
	if err != nil {
		fmt.Println("error while getting task : ", err)
		return
	}

	fmt.Println("getting task details : ", task.GetTaskID(), task.GetTitle(), task.GetDescription(), task.GetCreatedAt(), task.GetCreatedBy(), task.GetTaskType().String())

	//=========================================================================================================
	// create a sprint
	//NewCreateSprintDto(name string, startdate string, enddate string, createdBy entities.User)
	sprint, err := sprintController.CreateSprint(dtos.NewCreateSprintDto("Sprint 1", "2021-01-01", "2021-01-15", *user))
	if err != nil {
		fmt.Println("error while creating sprint : ", err)
		return
	}

	// add task to sprint
	_, err = sprintController.AddTaskToSprint(dtos.NewAddTaskToSprintDto(sprint.GetSprintID(), task.GetTaskID()))
	if err != nil {
		fmt.Println("error while adding task to sprint : ", err)
		return
	}

	//=========================================================================================================
	// get sprint

	getSprintDto := dtos.NewGetSprintDto(sprint.GetSprintID())
	sprint, err = sprintController.GetSprint(getSprintDto)
	if err != nil {
		fmt.Println("error while getting sprint : ", err)
		return
	}

	fmt.Println("getting sprint details : ", sprint.GetSprintID(), sprint.GetName(), sprint.GetStartDate(), sprint.GetEndDate())

	//=========================================================================================================
	// change the status of task
	_, err = taskController.ChangeTaskStatus(dtos.NewChangeTaskStatusDto(task.GetTaskID(), entities.IN_PROGRESS))
	if err != nil {
		fmt.Println("error while changing task status : ", err)
		return
	}

	//=========================================================================================================
	// get the status of task which we have changed recently
	task, err = taskController.GetTask(getTaskDto)
	if err != nil {
		fmt.Println("error while getting task : ", err)
		return
	}

	fmt.Println("getting task status on change : ", task.GetStatus().String())

	//=========================================================================================================
	// print all the tasks assigned to a user
	tasks, err := taskController.GetTasksByAssignee(dtos.NewGetTasksByAssigneeDto(assigneeUser.GetUserID()))
	if err != nil {
		fmt.Println("error while getting tasks by assignee : ", err)
		return
	}

	fmt.Println("Below tasks assigned to user : ", assigneeUser.GetName())
	for idx, assignedTask := range tasks {
		fmt.Println("Task ", idx, " : ", assignedTask.GetTitle(), assignedTask.GetDescription())
	}

	//=========================================================================================================

}
