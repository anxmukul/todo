package controller

import (
	"github.com/anxmukul/todo/model"
	"github.com/anxmukul/todo/view"
)

type Controller interface {
	Create(string, string) (*model.ToDo, error)
	SearchById(int64) (*model.ToDo, error)
	SearchByTitle(string) (*[]model.ToDo, error)
	DeleteByTitle(string) (*model.ToDo, error)
}

type TodoController struct {
	model model.TodoModel // why model is satisfies modelInterface, why it's not struct var?
}

func (t TodoController) Create(title string, content string) (*model.ToDo, error) {
	// fmt.Printf("In Controller %s\t%s\n", title, content)
	model := model.DbBasedModel()
	newTodo, err := model.CreateTodo(title, content)
	if err != nil {
		return nil, err
	}
	todoView := view.GetTodoView()
	todoView.ShowTodo(newTodo.Id, newTodo.Title, newTodo.Content)
	return newTodo, err

}

func (t TodoController) SearchById(id int64) (*model.ToDo, error) {
	model := model.DbBasedModel()
	todo, err := model.GetTodoById(id)
	if err != nil {
		return nil, err
	}
	todoView := view.GetTodoView()
	todoView.ShowTodo(todo.Id, todo.Title, todo.Content)
	return todo, err

}

func (t TodoController) SearchByTitle(title string) (*[]model.ToDo, error) {
	model := model.DbBasedModel()
	todos, err := model.GetTodoByTitle(title)
	if err != nil {
		return nil, err
	}
	todoView := view.GetTodoView()
	for _, element := range *todos {
		todoView.ShowTodo(element.Id, element.Title, element.Content)
	}
	return todos, err

}

func (t TodoController) DeleteByTitle(title string) (*model.ToDo, error) {
	model := model.DbBasedModel()
	todo, err := model.DeleteByTitle(title)
	if err != nil {
		return nil, err
	}
	todoView := view.GetTodoView()
	todoView.ShowTodo(todo.Id, todo.Title, todo.Content)
	return todo, err
}

func NewTodoController() Controller {
	return TodoController{
		model: model.DbBasedModel(), // how this todo{} struct compatible with Controller Interface return type.
	}
}
