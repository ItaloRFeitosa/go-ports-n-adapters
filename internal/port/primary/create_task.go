package primary

type CreateTaskInput struct {
	Description string `json:"description"`
}

type CreateTaskResult struct {
	ID string `json:"id"`
}
