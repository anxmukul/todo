package controller

import (
	"github.com/anxmukul/todo/model"
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
	newTodo , err := model.CreateTodo(title, content);

	return newTodo, err

}

func (t TodoController) SearchById(id int64) (*model.ToDo, error) {
	model := model.DbBasedModel()
	todo, err := model.GetTodoById(id)
	return todo, err;

}

func (t TodoController) SearchByTitle(title string) (*[]model.ToDo, error) {
	model := model.DbBasedModel()
	todos, err := model.GetTodoByTitle(title)
	return todos, err;

}

func (t TodoController) DeleteByTitle(string) (*model.ToDo, error) {
	return nil, nil
}

func NewTodoController() Controller {
	return TodoController{
		model : model.DbBasedModel(),	// how this todo{} struct compatible with Controller Interface return type.
	}
}	