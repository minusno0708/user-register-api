package model

import (
	"database/sql"
	
	_ "github.com/go-sql-driver/mysql"
)

var dbAddress = "db:3306"

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root@tcp("+dbAddress+")/test_db")
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func InsertUser(db *sql.DB, user User) error {
	_, err := db.Exec(
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

func findUserByUserID(db *sql.DB, userID string) (User, error) {
	var user User
	err := db.QueryRow(
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

func UpdateUsername(db *sql.DB, userID string, username string) error {
	_, err := db.Exec(
		"UPDATE users SET username = ? WHERE user_id = ?",
		username,
		userID,
	)
	if err != nil {
		return err
	}

	return nil
}

func DeleteUser(db *sql.DB, userID string) error {
	_, err := db.Exec(
		"DELETE FROM users WHERE user_id = ?",
		userID,
	)
	if err != nil {
		return err
	}

	return nil
}