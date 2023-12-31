package main

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"testing"
)

const endpoint = "http://localhost:8080"

var response struct {
	Message string `json:"message"`
}

func TestConnectionApi(t *testing.T) {
	expectedMessage := "Connection Successful"

	resp, err := http.Get(endpoint)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}
	responseData, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(responseData, &response)
	if err != nil {
		t.Fatal(err)
	}

  	if response.Message != expectedMessage {
		t.Fatalf("Expected message %v, got %v", expectedMessage, response.Message)
	}
}

func TestSigninBodyNotExist(t *testing.T) {
	expectedMessage := "Body is not valid"

	req, err := http.NewRequest("POST", endpoint+"/signin", nil)
	if err != nil {
		t.Fatal(err)
	}
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf("Expected status code 400, got %v", resp.StatusCode)
	}

	responseData, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(responseData, &response)
	if err != nil {
		t.Fatal(err)
	}

	if response.Message != expectedMessage {
		t.Fatalf("Expected message %v, got %v", expectedMessage, response.Message)
	}
	
}