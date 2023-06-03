package primary

type CreateTodoInput struct {
	Description string `json:"description"`
}

type CreateTodoResult struct {
	ID string `json:"id"`
}
