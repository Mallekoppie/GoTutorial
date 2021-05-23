package main

import (
	"database/sql"
	"log"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/zap"
)

func main() {

	dbSqlite, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Println("Error opening database", zap.Error(err))
		return
	}

	statement, _ :=
		dbSqlite.Prepare("CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)")
	statement.Exec()
	statement, _ =
		dbSqlite.Prepare("INSERT INTO people (firstname, lastname) VALUES (?, ?)")
	statement.Exec("Rob", "Gronkowski")
	rows, _ :=
		dbSqlite.Query("SELECT id, firstname, lastname FROM people")
	var id int
	var firstname string
	var lastname string
	for rows.Next() {
		rows.Scan(&id, &firstname, &lastname)
		log.Println(strconv.Itoa(id) + ": " + firstname + " " + lastname)
	}

}
