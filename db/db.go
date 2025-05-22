// db/db.go
package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/denisenkom/go-mssqldb"
)

var DB *sql.DB

func Init() {
	var err error
	dsn := os.Getenv("DB_DSN")
	DB, err = sql.Open("sqlserver", dsn)
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}
	err = DB.Ping()
	if err != nil {
		log.Fatal("Error connecting to DB: ", err)
	}
}
