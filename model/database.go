package model

import (
	"database/sql"
	
	_ "github.com/go-sql-driver/mysql"
)

var dbAddress = "0.0.0.0:3306"

func connectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root@tcp("+dbAddress+")/test_db")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func InsertUser(user User) error {
	db, err := connectDB()
	if err != nil {
		return err
	}
	defer db.Close()
	
	_, err = db.Exec(
		"INSERT INTO users (user_id, username, password) VALUES (?, ?, ?)",
		user.UserID,
		user.Username,
		user.Password,
	)
	if err != nil {
		return err
	}

	return nil
}

func findUserByUserID(userID string) (User, error) {
	db, err := connectDB()
	if err != nil {
		return User{}, err
	}
	defer db.Close()

	var user User
	err = db.QueryRow(
		"SELECT user_id, username, password FROM users WHERE user_id = ?",
		userID,
	).Scan(
		&user.UserID,
		&user.Username,
		&user.Password,
	)
	if err != nil {
		return User{}, err
	}

	return user, nil
}