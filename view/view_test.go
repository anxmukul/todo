package view

import (
	"testing"
)

func TestShowTodo(t *testing.T) {
	var todo Todo
	err := todo.ShowTodo(1, "title", "content")
	if err != nil {
		t.Errorf("got %s expected nil", err)
	}
}

func TestShowManyTodo(t *testing.T) {
	var todo Todo
	var arr []Todo
	err := todo.ShowManyTodo(&arr)
	if err != nil {
		t.Errorf("got %s expected nil", err)
	}
}
