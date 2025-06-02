package task

import (
	"fmt"
)

type Task struct {
	ID        int    `json:"id"`
	Tittle    string `json:"tittle"`
	Completed bool   `json:"completed"`
}

func addTask(p Task, id int, Tittle string, Completed bool) {

	var task string
	fmt.Println("Insert your Tittle: ")
	fmt.Scan(&task)

}
