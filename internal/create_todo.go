package internal

import (
	"github.com/google/uuid"
)

type CreateTodoInput struct {
	Description string `json:"description"`
}

type CreateTodoResult struct {
	ID          string `json:"id"`
	Description string `json:"description"`
}

func CreateTodo(req CreateTodoInput) (CreateTodoResult, error) {
	var res CreateTodoResult

	res.ID = uuid.NewString()
	res.Description = req.Description
	return res, nil
}
