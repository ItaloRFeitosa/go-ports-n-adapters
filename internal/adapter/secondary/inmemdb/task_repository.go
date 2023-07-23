package inmemdb

import (
	"errors"
	"fmt"
	"sync"

	"github.com/italorfeitosa/go-ports-n-adapters/internal/core/entity"
	secondaryport "github.com/italorfeitosa/go-ports-n-adapters/internal/port/secondary"
	"github.com/italorfeitosa/go-ports-n-adapters/pkg/uid"
)

var ErrTaskIDAlreadyTaken = errors.New("task id already taken")

type TaskRepository struct {
	sm sync.Map
}

// Ensure interface compliance
var _ secondaryport.TaskRepository = &TaskRepository{}

func NewTaskRepository() *TaskRepository {
	return &TaskRepository{sm: sync.Map{}}
}

func (r *TaskRepository) Insert(task *entity.Task) error {
	_, err := r.Get(task.ID)
	if err == nil {
		return ErrTaskIDAlreadyTaken
	}

	r.sm.Store(task.ID.String(), task)
	return nil
}

func (r *TaskRepository) Get(taskId uid.UID) (*entity.Task, error) {
	task, ok := r.sm.Load(taskId.String())
	if !ok {
		return nil, fmt.Errorf("not found")
	}

	return task.(*entity.Task), nil
}
