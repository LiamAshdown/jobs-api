package database

import (
	"database/sql"
	"fmt"
	"jobs-api/config"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func IntialiseDB() {
	fmt.Println("Initialising database")

	cfg := mysql.Config{
		User:                 config.GetConfig().Database.User,
		Passwd:               config.GetConfig().Database.Password,
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("%s:%d", config.GetConfig().Database.Host, config.GetConfig().Database.Port),
		DBName:               config.GetConfig().Database.DataBase,
		AllowNativePasswords: true,
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	// Open 10 simultaneous connections to the database
	// from my research seems database/sql handles connection pooling, so don't need to implement it myself
	db.SetMaxOpenConns(10)
	db.SetConnMaxIdleTime(5)
}

func GetDB() *sql.DB {
	return db
}
