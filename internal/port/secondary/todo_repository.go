package secondary

import "github.com/italorfeitosa/go-ports-n-adapters/internal/core/model"

//go:generate mockery --name TodoRepository
type TodoRepository interface {
	Insert(model.Todo) error
}
