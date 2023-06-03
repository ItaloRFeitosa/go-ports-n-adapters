package feature

import (
	"time"

	"github.com/google/uuid"
	"github.com/italorfeitosa/go-ports-n-adapters/internal/core/model"
	primaryport "github.com/italorfeitosa/go-ports-n-adapters/internal/port/primary"
	secondaryport "github.com/italorfeitosa/go-ports-n-adapters/internal/port/secondary"
)

type CreateTodoFeature struct {
	todoRepo secondaryport.TodoRepository
}

func NewCreateTodoFeature(todoRepo secondaryport.TodoRepository) *CreateTodoFeature {
	return &CreateTodoFeature{todoRepo}
}

func (f *CreateTodoFeature) CreateTodo(input primaryport.CreateTodoInput) (primaryport.CreateTodoResult, error) {
	var (
		todo   model.Todo
		result primaryport.CreateTodoResult
	)

	todo.ID = uuid.NewString()
	todo.Description = input.Description
	todo.CreatedAt = time.Now()

	if err := f.todoRepo.Insert(todo); err != nil {
		return result, err
	}

	result.ID = todo.ID

	return result, nil
}
