package handler

import (
	"todo-app/application/usecase"

	"github.com/gin-gonic/gin"
)

func GetTask(c *gin.Context) {
	usecase.GetTask(c)
}
