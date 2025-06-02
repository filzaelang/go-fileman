package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/denisenkom/go-mssqldb"
)

var DB_DEV *sql.DB

func Init_Dev() {
	var err error
	dsn := os.Getenv("DB_DSN_DEV")
	DB_DEV, err = sql.Open("sqlserver", dsn)
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}
	err = DB_DEV.Ping()
	if err != nil {
		log.Fatal("Error connecting to DB: ", err)
	}
}
