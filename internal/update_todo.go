package internal

type UpdateTodoInput struct {
	ID             string `uri:"id" binding:"required,uuid"`
	NewDescription string `json:"newDescription"`
}

type UpdateTodoResult struct {
	ID          string `json:"id"`
	Description string `json:"description"`
}

func UpdateTodo(req UpdateTodoInput) (UpdateTodoResult, error) {
	var todo UpdateTodoResult
	todo.ID = req.ID
	todo.Description = req.NewDescription
	return todo, nil
}
