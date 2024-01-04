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
	db, err := ConnectDB()
	if err != nil {
		t.Fatalf("Error connecting to the database: %v", err)
	}
	defer db.Close()
}

func TestInsertUser(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	err = InsertUser(db, testUser)
	if err != nil {
		t.Error(err)
	}
}

func TestInsertUserDuplicate(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	err = InsertUser(db, testUser)
	if err == nil {
		t.Error("Expected error, got nil")
	}
}

func TestFindUserByUserID(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	user, err := findUserByUserID(db, testUser.UserID)
	if err != nil {
		t.Error(err)
	}
	if user.UserID != testUser.UserID {
		t.Errorf("UserID is not match")
	}
	if user.Username != testUser.Username {
		t.Errorf("Username is not match")
	}
	if user.Password != testUser.Password {
		t.Errorf("Password is not match")
	}
}

func TestUpdateUsername(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	updatedName := "testuser_db_updated"
	err = UpdateUsername(db, testUser.UserID, updatedName)
	if err != nil {
		t.Error(err)
	}

	user, err := findUserByUserID(db, testUser.UserID)
	if err != nil {
		t.Error(err)
	}
	if user.Username != updatedName {
		t.Errorf("Failed to update username")
	}
}

func TestDeleteUser(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	err = DeleteUser(db, testUser.UserID)
	if err != nil {
		t.Error(err)
	}
}