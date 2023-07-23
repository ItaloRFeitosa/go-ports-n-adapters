package primary

import (
	"time"
)

type CreateTaskInput struct {
	Description string
}

type CreateTaskResult struct {
	ID string
}

type GetTaskResult struct {
	ID          string
	Description string
	DoneAt      *time.Time
	CreatedAt   time.Time
}
