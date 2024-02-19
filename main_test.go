package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// function to test the GET api
func TestTasksHandler(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(taskHandler))
	resp, err := http.Get(server.URL)
	statusCode := resp.StatusCode

	fmt.Println("The status code received is ", statusCode)
	fmt.Println("The value of err is ", err)

	if statusCode != http.StatusOK {
		t.Errorf("Expected status-code: 200 but received %v", resp.StatusCode)
	}
}
