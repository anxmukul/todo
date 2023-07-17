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
}

func (t ToDo) CreateTodo(title string, content string) (*ToDo, error) {
	connectToDb()
	insertQuery := `insert into mytodo(title, content) values($1, $2) returning id;`
	var id int
	err := db.QueryRow(insertQuery, title, content).Scan(&id)
	if err != nil {
		return nil, err
	}
	t = ToDo{Id: int64(id), Title: title, Content: content}

	return &t, nil
}

func (t ToDo) GetTodoById(id int64) (*ToDo, error) {
	connectToDb()
	searchQuery := `select * from mytodo where id = $1;`
	err := db.QueryRow(searchQuery, id).Scan(&t.Id, &t.Title, &t.Content)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (t ToDo) GetTodoByTitle(title string) (*[]ToDo, error) {
	connectToDb()
	searchQuery := `select * from mytodo where title = $1;`
	rows, err := db.Query(searchQuery, title)
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
	connectToDb()
	deleteQuery := `delete from mytodo where title = $1 returning id`
	err := db.QueryRow(deleteQuery, title).Scan(&t.Id)
	if err != nil {
		fmt.Println(err)
	}
	return &t, err
}

func DbBasedModel() TodoModel {
	return ToDo{}
}
