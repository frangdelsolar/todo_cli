package tests

import (
	"testing"
	"todo_cli/data"
)

func TestCreateTask(t *testing.T) {
	title := "Test Task"
	task, err := data.CreateTask(title)
	if err != nil {
		t.Errorf("CreateTask() returned an error: %v", err)
	}

	if task.Title != title {
		t.Errorf("CreateTask() returned the wrong task. Expected: %s, got: %s", title, task.Title)
	}
}

func TestGetTaskById(t *testing.T) {
	task, err := data.GetTaskById(1)
	if err != nil {
		t.Errorf("GetTaskById() returned an error: %v", err)
	}

	if task.Title != "Test Task" {
		t.Errorf("GetTaskById() returned the wrong task. Expected: %s, got: %s", "Test Task", task.Title)
	}
}

func TestUpdateTask(t *testing.T) {
	task := data.GetAllTasks()[0]
	updatedTitle := "Some test title"
	task.Title = updatedTitle
	task, err := data.UpdateTask(task.ID, task.Title)
	if err != nil {
		t.Errorf("UpdateTask() returned an error: %v", err)
	}
	if task.Title != updatedTitle {
		t.Errorf("UpdateTask() returned the wrong task. Expected: %s, got: %s", "Updated Task", task.Title)
	}
}

func TestGetAllTasks(t *testing.T) {
	tasks := data.GetAllTasks()
	if len(tasks) != 1 {
		t.Errorf("GetAllTasks() returned the wrong number of tasks. Expected: %d, got: %d", 1, len(tasks))
	}

	_, err := data.CreateTask("Second Task")
	if err != nil {
		t.Errorf("CreateTask() returned an error: %v", err)
	}

	tasks = data.GetAllTasks()
	if len(tasks) != 2 {
		t.Errorf("GetAllTasks() returned the wrong number of tasks. Expected: %d, got: %d", 2, len(tasks))
	}
}

func TestDeleteTask(t *testing.T) {
	task := data.GetAllTasks()[0]
	err := data.DeleteTask(task.ID)
	if err != nil {
		t.Errorf("DeleteTask() returned an error: %v", err)
	}
}
