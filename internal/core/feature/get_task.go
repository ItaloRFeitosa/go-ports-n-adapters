package feature

import (
	primaryport "github.com/italorfeitosa/go-ports-n-adapters/internal/port/primary"
	secondaryport "github.com/italorfeitosa/go-ports-n-adapters/internal/port/secondary"
	"github.com/italorfeitosa/go-ports-n-adapters/pkg/uid"
)

type getTaskFeature struct {
	taskRepository secondaryport.TaskRepository
}

func NewGetTaskFeature(taskRepository secondaryport.TaskRepository) *getTaskFeature {
	return &getTaskFeature{taskRepository}
}

func (f *getTaskFeature) GetTask(id string) (primaryport.GetTaskResult, error) {
	var (
		err     error
		taskDto primaryport.GetTaskResult
	)

	uid, err := uid.New(id)
	if err != nil {
		return taskDto, err
	}

	task, err := f.taskRepository.Get(uid)
	if err != nil {
		return taskDto, err
	}

	taskDto.ID = task.ID.String()
	taskDto.Description = task.Description
	taskDto.DoneAt = task.DoneAt
	taskDto.CreatedAt = task.CreatedAt

	return taskDto, nil
}
