package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "test:test@tcp(localhost:3306)/test")
	if err != nil {
		log.Println("Error opening DB: ", err.Error())
		return
	}
	insertData(db)
	callInsertStoredProcedure(db)
	selectData(db)

}

func callInsertStoredProcedure(db *sql.DB) {
	call, err := db.Prepare("call insert_user(?,?)")
	if err != nil {
		log.Fatalln("Error preparing statement: ", err.Error())
	}

	result, err := call.Exec("from go name", "from go surname")
	if err != nil {
		log.Fatalln("Error calling stored procedure to insert: ", err.Error())
	}

	affected, err := result.RowsAffected()
	if err != nil {
		log.Fatalln("Error getting rows affectred: ", err.Error())
	}
	log.Println("Stored procedure insert rows affected: ", affected)
}

func insertData(db *sql.DB) {
	insert, err := db.Prepare(`
	insert into user(name, surname) 
	values(?,?);`)
	if err != nil {
		log.Fatalln("Error preparing statement: ", err.Error())
	}

	result, err := insert.Exec("new inserted name", "new inserted surname")
	if err != nil {
		log.Fatalln("Error executing insert: ", err.Error())
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatalln("Error getting rows affected:", err.Error())
	}

	log.Println("Rows affected: ", rowsAffected)
}

func selectData(db *sql.DB) {
	resultset, err := db.Query("select * from user")
	if err != nil {
		log.Println("Error querying db:", err.Error())
		return
	}
	defer resultset.Close()

	for resultset.Next() {
		var id int
		var name, surname string
		err = resultset.Scan(&id, &name, &surname)
		if err != nil {
			log.Fatalln("Error getting data:", err.Error())
		}

		log.Printf("Retrieved data. Name %s Surname %s \n", name, surname)
	}
}
