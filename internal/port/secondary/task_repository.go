package secondary

import (
	"github.com/italorfeitosa/go-ports-n-adapters/internal/core/entity"
	"github.com/italorfeitosa/go-ports-n-adapters/pkg/uid"
)

type TaskRepository interface {
	Insert(*entity.Task) error
	Get(taskId uid.UID) (*entity.Task, error)
}
