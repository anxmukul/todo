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
	view  view.TodoDisplayer
}

func (t TodoController) Create(title string, content string) (*model.ToDo, error) {
	newTodo, err := t.model.CreateTodo(title, content)
	if err != nil {
		return nil, err
	}
	//newTodo: {id: 100, title: "foo"
	err = t.view.ShowTodo(newTodo.Id, newTodo.Title, newTodo.Content)
	if err != nil {
		return newTodo, err
	}
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
	err = t.view.ShowTodo(todo.Id, todo.Title, todo.Content)
	if err != nil {
		return todo, err
	}
	return todo, err

}

/*
1. given id
	a. it should return us todo and nil error when GetTodoByTitle return todo with repective id
	and nil err and ShowTodo return nil err
	b. it should return nil todo struct and non nil error wnen GetTodoByTitle fail and return nil todo and err.
	c. it should return todo struct and non nil error when GetTodoByTitle works fine but ShowTodo return non nil error
*/

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
	if err != nil {
		return todos, err
	}

	return todos, err

}

/*
Given title it should
	a. Return a todostruct  array and nil error
	b. it should return nil todo and non nil error when GetTodoByTitle return error
	c. It should return non nil todos array when model return todos array and non nil error, but Show Many todo return error
*/

func (t TodoController) DeleteByTitle(title string) (*model.ToDo, error) {
	todo, err := t.model.DeleteByTitle(title)
	if err != nil {
		return nil, err
	}
	err = t.view.ShowTodo(todo.Id, todo.Title, todo.Content)
	if err != nil {
		return todo, err
	}
	return todo, err
}

func NewTodoController() Controller {
	return TodoController{
		model: model.NewDbBasedModel(), // how this todo{} struct compatible with Controller Interface return type.
		view:  view.NewGetTodoView(),
	}
}
