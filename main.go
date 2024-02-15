package main

import (
	"fmt"
	"net/http"
)

// task represent a to-do task.
type Task struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

var tasks = []Task{}

const Dport = ":8012"

func main() {
	//	http.HandleFunc("/task", taskHandler)
	http.HandleFunc("/tasks", tasksHandler)
	fmt.Printf("Server is starting on port: %v", Dport)
	http.ListenAndServe(Dport, nil)
}

func taskHandler() {

}

func tasksHandler(w http.ResponseWriter, r *http.Request) {

}
