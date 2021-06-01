package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// ConnectDB : ...
func ConnectDB() *sql.DB {

	///////public
	dbUser := "ohm"
	dbPass := "ohmohm"
	dbName := "test_ohm"
	dbURL := "db.ondigitalocean.com:25060"

	var connectDB string = "" + dbUser + ":" + dbPass + "@tcp(" + dbURL + ")/" + dbName + ""

	conn, err := sql.Open("mysql", connectDB)
	if err != nil {
		fmt.Println("DB Connect Init::Err::", err)
	}

	return conn
}
