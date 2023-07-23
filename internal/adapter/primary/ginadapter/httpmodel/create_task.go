package httpmodel

import (
	"github.com/italorfeitosa/go-ports-n-adapters/internal/port/primary"
)

type CreateTaskRequest struct {
	Description string `json:"description"`
}

func (c CreateTaskRequest) ToInput() primary.CreateTaskInput {
	return primary.CreateTaskInput{
		Description: c.Description,
	}
}

type CreateTaskResponse struct {
	StatusCreated
	ID string `json:"id"`
}

func (CreateTaskResponse) FromOutput(result primary.CreateTaskResult) any {
	return CreateTaskResponse{
		ID: result.ID,
	}
}
