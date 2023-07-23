package httpmodel

import (
	"time"

	primaryport "github.com/italorfeitosa/go-ports-n-adapters/internal/port/primary"
)

type GetTaskRequest struct {
	ID string `uri:"id"`
}

func (c GetTaskRequest) ToInput() string {
	return c.ID
}

type GetTaskResponse struct {
	StatusOK
	ID          string     `json:"id"`
	Description string     `json:"description"`
	DoneAt      *time.Time `json:"doneAt"`
	CreatedAt   time.Time  `json:"createdAt"`
}

func (GetTaskResponse) FromOutput(result primaryport.GetTaskResult) any {
	return GetTaskResponse{
		ID:          result.ID,
		Description: result.Description,
		DoneAt:      result.DoneAt,
		CreatedAt:   result.CreatedAt,
	}
}
