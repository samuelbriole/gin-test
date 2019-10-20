package main

import "github.com/gin-gonic/gin"

func createRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	return router
}

func main() {
	router := createRouter()
	router.Run() // listen and serve on 0.0.0.0:8080
}
