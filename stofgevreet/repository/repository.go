package repository

import (
	"tutorial/stofgevreet/model"

	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// I know this is bad
func getDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "test:test@tcp(localhost:3306)/stofgevreet")
	if err != nil {
		log.Println("Error opening DB: ", err.Error())
		return nil, err
	}

	return db, nil
}

func SavePoint(input model.Point) error {
	db, err := getDB()
	if err != nil {
		return err
	}

	call, err := db.Prepare("call save_point(?,?,?,?,?,?)")
	if err != nil {
		log.Fatalln("Error preparing statement: ", err.Error())
	}

	result, err := call.Exec(input.Car, input.Scantime, input.Method,
		input.User, input.Points, input.Checkpoint)
	if err != nil {
		log.Fatalln("Error calling stored procedure to insert: ", err.Error())
	}

	affected, err := result.RowsAffected()
	if err != nil {
		log.Fatalln("Error getting rows affectred: ", err.Error())
	}
	log.Println("Stored procedure insert rows affected: ", affected)

	return nil
}

func SaveScan(input model.Scan) error {
	db, err := getDB()
	if err != nil {
		return err
	}

	call, err := db.Prepare("call save_scan(?,?,?,?)")
	if err != nil {
		log.Fatalln("Error preparing statement: ", err.Error())
	}

	result, err := call.Exec(input.Car, input.Scantime, input.Method, input.User)
	if err != nil {
		log.Fatalln("Error calling stored procedure to insert: ", err.Error())
	}

	affected, err := result.RowsAffected()
	if err != nil {
		log.Fatalln("Error getting rows affectred: ", err.Error())
	}
	log.Println("Stored procedure insert rows affected: ", affected)

	return nil
}

func SaveStopwatch(input model.Stopwatch) error {
	db, err := getDB()
	if err != nil {
		return err
	}

	call, err := db.Prepare("call save_stopwatch(?,?,?,?,?)")
	if err != nil {
		log.Fatalln("Error preparing statement: ", err.Error())
	}

	result, err := call.Exec(input.Car, input.Scantime, input.Method,
		input.User, input.Lap)
	if err != nil {
		log.Fatalln("Error calling stored procedure to insert: ", err.Error())
	}

	affected, err := result.RowsAffected()
	if err != nil {
		log.Fatalln("Error getting rows affectred: ", err.Error())
	}
	log.Println("Stored procedure insert rows affected: ", affected)

	return nil
}
