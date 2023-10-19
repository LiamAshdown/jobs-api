package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func IntialiseDB() {
	fmt.Println("Initialising database")
	dsn := "root:carbon12@tcp(localhost:3306)/jobs"

	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("Error connecting to the database: ", err)
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Error connecting to the database: ", err)
		panic(err)
	}
}

func GetDB() *sql.DB {
	return db
}
