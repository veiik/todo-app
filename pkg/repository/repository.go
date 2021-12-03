package repository

type TodoList interface {
}

type TodoItem interface {
}

type Authorisation interface {
}

type Repository struct {
	TodoItem
	TodoList
	Authorisation
}

func NewRepository() *Repository {
	return &Repository{}
}
