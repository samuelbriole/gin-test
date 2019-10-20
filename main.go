package main

import "github.com/samuelbriole/gin-test/presentation"

func main() {
	router := presentation.CreateRouter()
	router.Run("localhost:8080")
}
