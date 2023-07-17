package model

import (
	"fmt"
)

type TodoModel interface {
	CreateTodo(string, string) (*ToDo, error)
	GetTodoById(int64) (*ToDo, error)
	GetTodoByTitle(string) (*[]ToDo, error)
	DeleteByTitle(string) (*ToDo, error)
}

type ToDo struct {
	Id      int64
	Title   string
	Content string
	db      PostgresDB
}

func (t ToDo) CreateTodo(title string, content string) (*ToDo, error) {
	id, err := t.db.InsertTodo(title, content)
	if err != nil {
		return nil, err
	}
	t = ToDo{Id: int64(id), Title: title, Content: content}

	return &t, nil
}

func (t ToDo) GetTodoById(id int64) (*ToDo, error) {
	row := t.db.SearchTodoById(id)
	row.Scan(&t.Id, &t.Title, &t.Content)
	return &t, nil
}

func (t ToDo) GetTodoByTitle(title string) (*[]ToDo, error) {
	rows, err := t.db.SearchTodoByTitle(title)
	if err != nil {
		return nil, err
	}
	var todoArray = make([]ToDo, 0)
	for rows.Next() {
		err := rows.Scan(&t.Id, &t.Title, &t.Content)
		if err != nil {
			fmt.Println(err)
		}
		todoArray = append(todoArray, t)
	}
	return &todoArray, nil
}

func (t ToDo) DeleteByTitle(title string) (*ToDo, error) {
	id, err := t.db.DeleteTodo(title)
	if err != nil {
		fmt.Println(err)
	}
	t.Id = id
	return &t, err
}

func NewDbBasedModel() TodoModel {
	return ToDo{
		db: NewConnectToDb(),
	}
}
