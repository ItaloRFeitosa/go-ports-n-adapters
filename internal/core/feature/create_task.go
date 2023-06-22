package feature

import (
	"time"

	"github.com/italorfeitosa/go-ports-n-adapters/internal/core/model"
	primaryport "github.com/italorfeitosa/go-ports-n-adapters/internal/port/primary"
	secondaryport "github.com/italorfeitosa/go-ports-n-adapters/internal/port/secondary"
	"github.com/italorfeitosa/go-ports-n-adapters/pkg/uid"
)

type CreateTaskFeature struct {
	taskRepo secondaryport.TaskRepository
}

func NewCreateTaskFeature(todoRepo secondaryport.TaskRepository) *CreateTaskFeature {
	return &CreateTaskFeature{todoRepo}
}

func (f *CreateTaskFeature) CreateTask(input primaryport.CreateTaskInput) (primaryport.CreateTaskResult, error) {
	var (
		newTask model.Task
		result  primaryport.CreateTaskResult
	)

	newTask.ID = uid.Random()
	newTask.Description = input.Description
	newTask.CreatedAt = time.Now()

	if err := f.taskRepo.Insert(newTask); err != nil {
		return result, err
	}

	result.ID = newTask.ID.String()

	return result, nil
}
