package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func createRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello world")
	})
	return router
}

func main() {
	router := createRouter()
	router.Run() // listen and serve on 0.0.0.0:8080
}
