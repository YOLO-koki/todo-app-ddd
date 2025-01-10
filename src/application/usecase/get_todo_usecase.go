package usecase

import (
	"todo-app/infrastructure/repository"

	"github.com/gin-gonic/gin"
)

func GetTodo(c *gin.Context) {
	todo := repository.GetById(1)

	c.JSON(200, gin.H{
		"todo": todo,
	})
}
