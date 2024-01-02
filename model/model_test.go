package model

import (
	"testing"
)

func TestConnectionDB(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Error("Error connecting to database")
	}
	defer db.Close()
}