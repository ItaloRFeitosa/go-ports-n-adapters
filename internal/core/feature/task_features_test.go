package feature_test

import (
	"testing"

	"github.com/italorfeitosa/go-ports-n-adapters/internal/adapter/secondary/inmemdb"
	"github.com/italorfeitosa/go-ports-n-adapters/internal/core/feature"
	primaryport "github.com/italorfeitosa/go-ports-n-adapters/internal/port/primary"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateTask_GetTask(t *testing.T) {
	taskRepo := inmemdb.NewTaskRepository()
	createTaskFeature := feature.NewCreateTaskFeature(taskRepo)
	getTaskFeature := feature.NewGetTaskFeature(taskRepo)

	input := primaryport.CreateTaskInput{
		Description: "some description",
	}

	out, err := createTaskFeature.CreateTask(input)

	require.NoError(t, err)

	taskDto, err := getTaskFeature.GetTask(out.ID)

	require.NoError(t, err)

	assert.Equal(t, out.ID, taskDto.ID)
	assert.Equal(t, input.Description, taskDto.Description)
	assert.NotZero(t, taskDto.CreatedAt)
	assert.Nil(t, taskDto.DoneAt)
}
