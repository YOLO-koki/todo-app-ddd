package handler

import (
	"todo-app/application/usecase"

	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	usecase.Hello(c)
}