package entity

import (
	"fmt"
	"time"

	"github.com/italorfeitosa/go-ports-n-adapters/pkg/uid"
)

type Task struct {
	ID          uid.UID
	Description string
	DoneAt      *time.Time
	CreatedAt   time.Time
}

func CreateTask(description string) (*Task, error) {
	task := new(Task)
	task.ID = uid.Random()
	task.Description = description
	task.CreatedAt = time.Now()

	return task, task.Validate()
}

func (t *Task) Validate() error {
	if t.Description == "" {
		return fmt.Errorf("description must be set")
	}

	return nil
}
