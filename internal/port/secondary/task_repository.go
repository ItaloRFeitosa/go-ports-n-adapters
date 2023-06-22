package secondary

import "github.com/italorfeitosa/go-ports-n-adapters/internal/core/model"

type TaskRepository interface {
	Insert(model.Task) error
}
