package view

import (
	"fmt"
	"reflect"
	"testing"
)

func TestShowTodo(t *testing.T) {
	var todo Todo
	err := todo.ShowTodo(1, "title", "content")
	if err != nil {
		t.Errorf("got %s expected nil", err)
	} else {
		fmt.Printf("got %+v expected %+v", err, nil)
	}
}

func TestShowManyTodo(t *testing.T) {
	var todo Todo
	var arr []Todo
	err := todo.ShowManyTodo(&arr)
	if err != nil {
		t.Errorf("got %s expected nil", err)
	} else {
		fmt.Printf("got %+v expected %+v", err, nil)
	}
}

func TestNewGetTodoView(t *testing.T) {
	var todo Todo
	res := NewGetTodoView()
	fmt.Println("res = ", reflect.TypeOf(res))
	if reflect.TypeOf(res) != reflect.TypeOf(todo) {
		t.Errorf("got res of type %+v but expected %+v type", res, todo)	
	} else {
		fmt.Printf("got res of %+v type and expected %+v type", res, todo)
	}
}
