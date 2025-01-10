package presentation

import "github.com/gin-gonic/gin"

func Router() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.GET("/hello", hello)
	}

	return router
}

func hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello",
	})
}
