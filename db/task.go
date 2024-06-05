package db

import (
	"time"

	"github.com/gofrs/uuid"
)

type Task struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	CreatedAt string    `json:"createdAt"`
}

func NewTask(title string) *Task {
	dbase := GetDB()

	now := time.Now().Format(time.RFC3339)
	task := &Task{
		ID:        uuid.Must(uuid.NewV4()),
		Title:     title,
		CreatedAt: now,
	}

	dbase.Tasks[task.ID.String()] = *task
	dbase.Save()
	return task
}

func UpdateTask(id string, title string) *Task {
	task := GetTaskById(id)
	task.Update(title)
	return task
}

func DeleteTask(id string) {
	dbase := GetDB()
	delete(dbase.Tasks, id)
	dbase.Save()
}

func (t *Task) Update(title string) *Task {
	dbase := GetDB()
	task := dbase.Tasks[t.ID.String()]
	task.Title = title
	dbase.Save()
	return t
}

func (t *Task) Delete(){
	dbase := GetDB()
	delete(dbase.Tasks, t.ID.String())
	dbase.Save()
}

func GetTaskById(id string) *Task {
	dbase := GetDB()
	task := dbase.Tasks[id]
	return &task
}

func GetAllTasks() []Task {
	dbase := GetDB()
	tasks := make([]Task, 0, len(dbase.Tasks))
	for _, task := range dbase.Tasks {
		tasks = append(tasks, task)
	}
	return tasks
}
