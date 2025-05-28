package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/denisenkom/go-mssqldb"
)

var DB_MIS *sql.DB

func Init_Mis() {
	var err error
	dsn_mis := os.Getenv("DB_DSN_MIS")
	DB_MIS, err = sql.Open("sqlserver", dsn_mis)
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}
	err = DB_MIS.Ping()
	if err != nil {
		log.Fatal("Error connecting to DB: ", err)
	}
}
