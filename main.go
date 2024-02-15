package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
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

	switch r.Method {
	case "GET":
		json.NewEncoder(w).Encode(tasks)
	case "POST":
		var task Task
		if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// go get  github.com/google/uuid

		task.ID = uuid.New().String()
		tasks = append(tasks, task)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(tasks)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}
