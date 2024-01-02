package model

import (
	"testing"
)

var testUser = User{
	UserID:   "testuser_db",
	Username: "testuser_db",
	Password: "test_password",
}

func TestConnectionDB(t *testing.T) {
	db, err := connectDB()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
}

func TestInsertUser(t *testing.T) {
	err := InsertUser(testUser)
	if err != nil {
		t.Error(err)
	}
}

func TestFindUserByUserID(t *testing.T) {
	user, err := findUserByUserID(testUser.UserID)
	if err != nil {
		t.Error(err)
	}
	if user.UserID != testUser.UserID {
		t.Errorf("Expected %s, got %s", testUser.UserID, user.UserID)
	}
	if user.Username != testUser.Username {
		t.Errorf("Expected %s, got %s", testUser.Username, user.Username)
	}
	if user.Password != testUser.Password {
		t.Errorf("Password is not match")
	}
}