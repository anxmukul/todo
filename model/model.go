package model

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

var (
	host     = os.Getenv("host")
	port     = os.Getenv("port")
	user     = os.Getenv("user")
	password = os.Getenv("password")
	dbname   = os.Getenv("dbname")
)

func connectToDb() {
	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	var err error
	db, err = sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
}

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

func (t ToDo) DeleteByTitle(string) (*ToDo, error) {
	return nil, nil
}

func DbBasedModel() TodoModel {
	return ToDo{}
}
