package config

import (
	"testing"
)

func TestConnectionDB(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Fatalf("Error connecting to the database: %v", err)
	}
	defer db.Close()
}