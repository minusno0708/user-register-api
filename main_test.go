package main

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"testing"
)

func TestConnectionApi(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}
	responseData, _ := ioutil.ReadAll(resp.Body)

	var data struct {
		Message string `json:"message"`
	}

	err = json.Unmarshal(responseData, &data)
	if err != nil {
		t.Fatal(err)
	}
	expectedMessage := "Connection Successful"

  	if data.Message != expectedMessage {
		t.Fatalf("Expected message %v, got %v", expectedMessage, data.Message)
	}
}