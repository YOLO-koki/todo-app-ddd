package usecase

import (
	"todo-app/infrastructure/repository"

	"github.com/gin-gonic/gin"
)

func GetTask(c *gin.Context) {
	task := repository.GetById(1)

	c.JSON(200, gin.H{
		"task": task,
	})
}
