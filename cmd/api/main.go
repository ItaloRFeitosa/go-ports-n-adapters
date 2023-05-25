package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/italorfeitosa/go-ports-n-adapters/internal"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/todos", internal.AdaptController(http.StatusCreated, internal.CreateTodo))
	r.PATCH("/todos/:id", internal.AdaptController(http.StatusOK, internal.UpdateTodo))
	r.Run()
}
