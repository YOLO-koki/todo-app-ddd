package task

import "log"

type status string

const (
	pending    status = "pending"
	inProgress        = "in_progress"
	completed         = "completed"
)

type Task struct {
	ID          int
	Title       string
	Description string
	Status      status
}

func New(title string, description string) *Task {
	return &Task{Title: title, Description: description, Status: pending}
}

func ReNew(id int, title string, description string, status status) *Task {
	return &Task{ID: id, Title: title, Description: description, Status: status}
}

func ReNewStatus(s string) status {
	if s != string(pending) && s != string(inProgress) && s != string(completed) {
		log.Fatal("status is must be pending or in_progress or completed.")
	}
	return status(s)
}
