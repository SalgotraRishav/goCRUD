package main

// task represent a to-do task.
type Task struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

var tasks = []Task{}

const Dport = ":8012"

func main() {

}
