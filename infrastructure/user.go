package infrastructure

import (
	"database/sql"

	"user-register-api/domain"
)

func InsertUser(db *sql.DB, user domain.User) error {
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

func findUserByUserID(db *sql.DB, userID string) (domain.User, error) {
	var user domain.User
	err := db.QueryRow(
		"SELECT user_id, username, password FROM users WHERE user_id = ?",
		userID,
	).Scan(
		&user.UserID,
		&user.Username,
		&user.Password,
	)
	if err != nil {
		return domain.User{}, err
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