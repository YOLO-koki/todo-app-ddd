package presentation

import (
	"todo-app/presentation/handler"

	"github.com/gin-gonic/gin"
)

type AllHandler struct {
	taskHandler handler.TaskHandler
}

func Router() *gin.Engine {
	router := gin.Default()
	allHandler := AllHandler{taskHandler: handler.TaskHandler{}}

	v1 := router.Group("/v1")
	{
		v1.GET("/tasks", allHandler.taskHandler.GetTasks)
		v1.GET("/tasks/:id", allHandler.taskHandler.GetTask)
	}

	return router
}
