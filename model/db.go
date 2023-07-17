package model

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type PostgresDB interface {
	InsertTodo(string, string) (int64, error)
	DeleteTodo(string) (int64, error)
	SearchTodoById(int64) *sql.Row
	SearchTodoByTitle(string) (*sql.Rows, error)
}

type Database struct {
	Db *sql.DB
}

var (
	host     = os.Getenv("host")
	port     = os.Getenv("port")
	user     = os.Getenv("user")
	password = os.Getenv("password")
	dbname   = os.Getenv("dbname")
)

func (db Database) InsertTodo(title string, content string) (int64, error) {
	insertQuery := `insert into mytodo(title, content) values($1, $2) returning id;`
	var id int64
	err := db.Db.QueryRow(insertQuery, title, content).Scan(&id)
	return id, err
}

func (db Database) DeleteTodo(title string) (int64, error) {
	deleteQuery := `delete from mytodo where title = $1 returning id`
	var id int64
	err := db.Db.QueryRow(deleteQuery, title).Scan(&id)
	return id, err
}

func (db Database) SearchTodoById(id int64) *sql.Row {
	fmt.Println("Inside db.go searchtodobyid();")
	searchQuery := `select * from mytodo where id = $1;`
	rows := db.Db.QueryRow(searchQuery, id)
	return rows
}

func (db Database) SearchTodoByTitle(title string) (*sql.Rows, error) {
	searchQuery := `select * from mytodo where title = $1;`
	rows, err := db.Db.Query(searchQuery, title)
	return rows, err
}

func connectToDb() *sql.DB {
	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	// fmt.Println("Inside connectToDb() ")
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println(err);
	}
	return db
}

func NewConnectToDb() PostgresDB {
	return Database{
		Db: connectToDb(),
	}
}
