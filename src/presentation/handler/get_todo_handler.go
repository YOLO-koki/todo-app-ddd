package handler

import (
	"todo-app/application/usecase"

	"github.com/gin-gonic/gin"
)

func GetTitle(c *gin.Context) {
	usecase.GetTodo(c)
}
