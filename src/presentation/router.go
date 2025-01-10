package presentation

import (
	"todo-app/presentation/handler"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.GET("/todo", handler.GetTask)
	}

	return router
}
