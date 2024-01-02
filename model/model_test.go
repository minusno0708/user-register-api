package model

import (
	"testing"
)

func TestConnectionDB(t *testing.T) {
	db, err := connectDB()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
}

func TestInsertUser(t *testing.T) {
	user := User{
		UserID:   "testuser_db",
		Username: "testuser_db",
		Password: "test_password",
	}
	err := InsertUser(user)
	if err != nil {
		t.Error(err)
	}
}