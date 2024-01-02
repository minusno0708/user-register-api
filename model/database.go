package model

import (
	"database/sql"
	
	_ "github.com/go-sql-driver/mysql"
)

var dbAddress = "localhost:3306"

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:password@tcp("+dbAddress+")/test_db")
	if err != nil {
		return nil, err
	}
	return db, nil
}