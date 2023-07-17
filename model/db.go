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
