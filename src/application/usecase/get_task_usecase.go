package usecase

import (
	"todo-app/domain/task"
	"todo-app/infrastructure/repository"
)

type GetTasKUsecase struct {
	taskRepository repository.TaskRepository
}

func (u *GetTasKUsecase) GetTask(id int) *task.Task {
	task := u.taskRepository.FindById(id)
	return task
}
