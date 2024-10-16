package initializers

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func OpenConnection() *sql.DB {
	db, _ := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/pj_golang")
	return db
}
