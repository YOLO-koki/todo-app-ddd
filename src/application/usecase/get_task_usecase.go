package usecase

import (
	"todo-app/domain/task"
	"todo-app/infrastructure/repository"
)

type GetTaskUsecase struct {
	taskrepository repository.TaskRepository
}

func (u *GetTaskUsecase) GetTask(id int) *task.Task {
	t := repository.NewTaskRepository()
	task := t.FindById(id)
	return task
}
