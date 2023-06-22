package model

import (
	"time"

	"github.com/italorfeitosa/go-ports-n-adapters/internal/shared/uid"
)

type Task struct {
	ID          uid.UID
	Description string
	DoneAt      *time.Time
	CreatedAt   time.Time
}
