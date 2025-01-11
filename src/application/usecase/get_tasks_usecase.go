package usecase

import (
	"todo-app/domain/task"
	"todo-app/infrastructure/repository"

	"github.com/gin-gonic/gin"
)

type GetTasksUsecase struct {
	taskRepository repository.TaskRepository
}

func (u *GetTasksUsecase) GetTasks(c *gin.Context) []*task.Task {
	t := repository.NewTaskRepository()
	tasks := t.FindAll()

	return tasks
}
