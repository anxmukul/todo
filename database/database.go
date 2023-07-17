// package database

// import (
// 	"database/sql"
// 	"os"

// 	"github.com/anxmukul/todo/model"
// )

// type TodoStore interface {
// 	InsertTodo(string) *model.ToDo
// 	SearchTodoById(string) *model.ToDo
// 	SearchTodoByTitle(string) *[]model.ToDo
// 	DeleteTodoByTitle(string) *model.ToDo
// }

// var (
// 	host     = os.Getenv("host")
// 	port     = os.Getenv("port")
// 	user     = os.Getenv("user")
// 	password = os.Getenv("password")
// 	dbname   = os.Getenv("dbname")
// )

// type Db struct {
// 	db *sql.DB
// }

// func (db Db) InsertTodo(string) *model.ToDo {

// }

// func (db Db) SearchTodoById(string) *model.ToDo {

// }

// func (db Db) SearchTodoByTitle(string) *[]model.ToDo {

// }

// func (db Db) DeleteTodoByTitle(string) *model.ToDo {

// }
