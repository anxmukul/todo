package controller

import (
	"fmt"

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
	view  view.TodoDisplayer
}

func add (a, b int) int{
	return a / b;
}

func (t TodoController) Create(title string, content string) (*model.ToDo, error) {
	// fmt.Printf("In Controller %s\t%s\n", title, content)
	newTodo, err := t.model.CreateTodo(title, content)
	if err != nil {
		return nil, err
	}
	err = t.view.ShowTodo(newTodo.Id, newTodo.Title, newTodo.Content)
	return newTodo, err

}

/*
unit test cases of create
1. given title and content, 
	a. it should return us a todo struct and nil error when createToDo return a new Todo
and shwoTodo return nil error
	b. it should return us a nil todo struct and error when createToDO return error
	c. it should return us todo struct and error when createToDo return newTodo and nul error
	 but showToDo return non-nil error
*/

func (t TodoController) SearchById(id int64) (*model.ToDo, error) {
	todo, err := t.model.GetTodoById(id)
	if err != nil {
		return nil, err
	}
	fmt.Println(todo)

	err = t.view.ShowTodo(todo.Id, todo.Title, todo.Content)
	return todo, err

}

func (t TodoController) SearchByTitle(title string) (*[]model.ToDo, error) {
	todos, err := t.model.GetTodoByTitle(title)
	if err != nil {
		return nil, err
	}
	var todoArray []view.Todo
	for _, element := range *todos {
		viewtodo := view.Todo{Id: element.Id, Title: element.Title, Content: element.Content}
		todoArray = append(todoArray, viewtodo)
	}
	err = t.view.ShowManyTodo(&todoArray)
	return todos, err

}

func (t TodoController) DeleteByTitle(title string) (*model.ToDo, error) {
	todo, err := t.model.DeleteByTitle(title)
	if err != nil {
		return nil, err
	}
	err = t.view.ShowTodo(todo.Id, todo.Title, todo.Content)
	return todo, err
}

func NewTodoController() Controller {
	return TodoController{
		model: model.NewDbBasedModel(), // how this todo{} struct compatible with Controller Interface return type.
		view:  view.NewGetTodoView(),
	}
}
