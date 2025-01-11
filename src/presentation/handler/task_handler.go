package handler

import (
	"strconv"
	"todo-app/application/usecase"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	getTasksUsecase usecase.GetTasksUsecase
	getTaskUsecase  usecase.GetTaskUsecase
}

func (h *TaskHandler) GetTasks(c *gin.Context) {
	tasks := h.getTasksUsecase.GetTasks(c)

	c.JSON(200, tasks)
}

func (h *TaskHandler) GetTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	task := h.getTaskUsecase.GetTask(id)

	c.JSON(200, task)
}
