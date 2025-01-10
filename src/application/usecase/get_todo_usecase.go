package usecase

import (
	"github.com/gin-gonic/gin"
)

func GetTodo(c *gin.Context) {
	todo := "todo"
	c.JSON(200, gin.H{
		"title": todo,
	})
}
