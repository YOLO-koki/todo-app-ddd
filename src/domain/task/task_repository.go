package task

type TaskRepository interface {
	FindById(id int) *Task
	FindAll() []*Task
}
