package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

var db *sql.DB

func connect() *sql.DB {

	/*define variables in bulk*/
	var (
		dbname = os.Getenv("DB_NAME")
		dbport = os.Getenv("DB_PORT")
		dbhost = os.Getenv("DB_HOST")
		dbpass = os.Getenv("DB_PASSWORD")
		dbuser = os.Getenv("DB_USER")
	)

	dburl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbuser, dbpass, dbhost, dbport, dbname)
	db, err := sql.Open("mysql", dburl)
	if err != nil {
		log.Fatal("Failed to connect")
	}
	fmt.Println("Connected")
	return db
}
