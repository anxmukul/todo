package model

type TodoModel interface {
	CreateTodo(string, string) (*ToDo, error)
	GetTodoById(int64) (*ToDo, error)
	GetTodoByTitle(string) (*[]ToDo, error)
	DeleteByTitle(string) (*ToDo, error)
}

type ToDo struct {
	Id int64
	Title string
	Content string
}

