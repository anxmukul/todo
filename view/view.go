package view

import (
	"encoding/json"
	"errors"
	"fmt"
)

type TodoDisplayer interface {
	ShowTodo(int64, string, string) error
	ShowManyTodo(*[]Todo) error
}

type Todo struct {
	Id      int64
	Title   string
	Content string
}

func (t Todo) ShowTodo(id int64, title string, content string) error {
	// fmt.Println("This will display one todo")
	if id <= 0 {
		err := errors.New("Empty: no todo available to displat")
		return err
	}
	t.Id = id
	t.Title = title
	t.Content = content
	todos, _ := json.MarshalIndent(t, "", "  ")
	fmt.Println(string(todos))
	return nil
}

func (t Todo) ShowManyTodo(todoArray *[]Todo) error {
	// fmt.Println("This will display Many todo")
	todos, _ := json.MarshalIndent(todoArray, "", "  ")
	fmt.Println(string(todos))
	return nil
}

func NewGetTodoView() TodoDisplayer {
	return Todo{}
}
