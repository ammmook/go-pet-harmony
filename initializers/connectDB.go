package initializers

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDatabase() {
	var err error
	DB, err = sql.Open("mysql", "golang:1234@tcp(localhost:33091)/pj_golang")
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Failed to ping the database: ", err)
	}

	defer DB.Close()

	fmt.Println("Connect Database Succesful!")

}
