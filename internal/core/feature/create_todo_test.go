package feature_test

import (
	"fmt"
	"testing"

	"github.com/italorfeitosa/go-ports-n-adapters/internal/core/feature"
	"github.com/italorfeitosa/go-ports-n-adapters/internal/core/model"
	primaryport "github.com/italorfeitosa/go-ports-n-adapters/internal/port/primary"
	"github.com/stretchr/testify/assert"
)

func TestCreateTodoFeature(t *testing.T) {

	tests := []struct {
		name     string
		todoRepo *TodoRepositoryMock
		input    primaryport.CreateTodoInput
		wantErr  bool
	}{
		{
			name:     "given a CreateTodoInput should call repo Insert and then return todo id",
			todoRepo: new(TodoRepositoryMock),
			input: primaryport.CreateTodoInput{
				Description: "some description",
			},
		},
		{
			name:     "given a CreateTodoInput when repo Insert return error should return this error",
			todoRepo: new(TodoRepositoryMock).WithError(fmt.Errorf("some error")),
			input: primaryport.CreateTodoInput{
				Description: "some description",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			createTodoFeature := feature.NewCreateTodoFeature(tt.todoRepo)

			result, err := createTodoFeature.CreateTodo(tt.input)

			assert.Equal(t, tt.todoRepo.InsertReceivedParam.Description, tt.input.Description)
			assert.Nil(t, tt.todoRepo.InsertReceivedParam.DoneAt)

			if tt.wantErr {
				assert.Equal(t, tt.todoRepo.Error, err)
			} else {
				assert.Equal(t, tt.todoRepo.InsertReceivedParam.ID, result.ID)
				assert.NoError(t, err)
			}

		})
	}
}

type TodoRepositoryMock struct {
	InsertReceivedParam model.Todo

	Error error
}

func (mock *TodoRepositoryMock) Insert(todo model.Todo) error {
	mock.InsertReceivedParam = todo

	return mock.Error
}

func (mock *TodoRepositoryMock) WithError(err error) *TodoRepositoryMock {
	mock.Error = err

	return mock
}
