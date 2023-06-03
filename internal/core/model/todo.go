package model

import "time"

type Todo struct {
	ID          string
	Description string
	DoneAt      *time.Time
	CreatedAt   time.Time
}
