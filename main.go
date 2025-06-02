package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Task struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Insert your Title: ")
	scanner.Scan()
	temp := scanner.Text()

	task := Task{Id: 1, Title: temp, Completed: false}

	a, err := json.Marshal(task)
	if err != nil {
		fmt.Println("Error marshaling JSON :", err)
		return
	}
	fmt.Println(string(a))

	filename := "Task_" + strconv.Itoa(task.Id) + ".json"
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating File: ", err)
		return
	}
	defer file.Close()

	_, err = file.Write(a)
	if err != nil {
		fmt.Println("Error Wrting FIle : ", err)
		return
	}
}
