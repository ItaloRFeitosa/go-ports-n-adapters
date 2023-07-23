package feature_test

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/italorfeitosa/go-ports-n-adapters/internal/core/entity"
	"github.com/italorfeitosa/go-ports-n-adapters/internal/core/feature"
	primaryport "github.com/italorfeitosa/go-ports-n-adapters/internal/port/primary"
	"github.com/italorfeitosa/go-ports-n-adapters/internal/port/secondary/mocks"
	"github.com/italorfeitosa/go-ports-n-adapters/pkg/uid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type testCase struct {
	name         string
	todoRepoMock *mocks.TaskRepository
	input        primaryport.CreateTaskInput
	prepare      func(tc *testCase)
	wantErr      error
}

func prepareTest(tc *testCase) {
	tc.todoRepoMock.EXPECT().Insert(mock.MatchedBy(func(todo *entity.Task) bool {
		return tc.input.Description == todo.Description &&
			todo.ID.Valid() && todo.DoneAt == nil && !todo.CreatedAt.IsZero()
	})).Return(tc.wantErr)
}

func TestCreateTodoFeature(t *testing.T) {

	cases := []testCase{
		{
			name:    "given a CreateTodoInput should call repo Insert and then return valid todo id",
			prepare: prepareTest,
			input: primaryport.CreateTaskInput{
				Description: "some description",
			},
		},
		{
			name:    "given a CreateTodoInput when repo Insert return error should return this error",
			prepare: prepareTest,
			input: primaryport.CreateTaskInput{
				Description: "some description",
			},
			wantErr: gofakeit.ErrorDatabase(),
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.todoRepoMock = mocks.NewTaskRepository(t)

			tc.prepare(&tc)

			feature := feature.NewCreateTaskFeature(tc.todoRepoMock)

			out, err := feature.CreateTask(tc.input)

			if tc.wantErr != nil {
				assert.Equal(t, tc.wantErr, err)
			} else {
				assert.True(t, uid.Validate(out.ID))
			}
		})
	}
}
