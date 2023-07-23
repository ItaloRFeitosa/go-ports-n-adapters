package feature

import (
	"github.com/italorfeitosa/go-ports-n-adapters/internal/core/entity"
	primaryport "github.com/italorfeitosa/go-ports-n-adapters/internal/port/primary"
	secondaryport "github.com/italorfeitosa/go-ports-n-adapters/internal/port/secondary"
)

type createTaskFeature struct {
	taskRepository secondaryport.TaskRepository
}

func NewCreateTaskFeature(taskRepository secondaryport.TaskRepository) *createTaskFeature {
	return &createTaskFeature{taskRepository}
}

func (f *createTaskFeature) CreateTask(input primaryport.CreateTaskInput) (primaryport.CreateTaskResult, error) {
	var result primaryport.CreateTaskResult

	newTask, err := entity.CreateTask(input.Description)
	if err != nil {
		return result, err
	}

	if err := f.taskRepository.Insert(newTask); err != nil {
		return result, err
	}

	result.ID = newTask.ID.String()

	return result, nil
}
