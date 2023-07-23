package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/italorfeitosa/go-ports-n-adapters/internal/adapter/primary/ginadapter"
	"github.com/italorfeitosa/go-ports-n-adapters/internal/adapter/primary/ginadapter/httpmodel"
	"github.com/italorfeitosa/go-ports-n-adapters/internal/adapter/secondary/inmemdb"
	"github.com/italorfeitosa/go-ports-n-adapters/internal/core/feature"
)

func main() {
	taskRepository := inmemdb.NewTaskRepository()
	createTaskFeature := feature.NewCreateTaskFeature(taskRepository)
	getTaskFeature := feature.NewGetTaskFeature(taskRepository)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/tasks", ginadapter.Adapt(httpmodel.CreateTaskRequest{}, createTaskFeature.CreateTask, httpmodel.CreateTaskResponse{}))
	r.GET("/tasks/:id", ginadapter.Adapt(httpmodel.GetTaskRequest{}, getTaskFeature.GetTask, httpmodel.GetTaskResponse{}))

	r.Run()
}
