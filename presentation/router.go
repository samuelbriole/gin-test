package presentation

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/handler"
)

func CreateRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello world")
	})
	router.Any("/graphql", func(c *gin.Context) {
		var graphqlHandler = handler.New(&handler.Config{
			Schema:   &Schema,
			Pretty:   true,
			GraphiQL: true,
		})
		graphqlHandler.ServeHTTP(c.Writer, c.Request)
	})
	return router
}
