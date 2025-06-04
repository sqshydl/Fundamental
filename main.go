package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

type autoInc struct {
	sync.Mutex
	id int
}

func (ai *autoInc) ID() (id int) {
	ai.Lock()
	defer ai.Unlock()

	id = ai.id
	ai.id++
	return
}

type Task struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	ai := autoInc{}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Insert your Title: ")
	scanner.Scan()
	temp := scanner.Text()

	task := Task{Id: ai.ID(), Title: temp, Completed: false}

	a, err := json.MarshalIndent(task, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON :", err)
		return
	}
	fmt.Println(string(a))

	os.MkdirAll("storage", os.ModePerm)

	// Open file for writing/appending
	file, err := os.OpenFile("storage/Task.json", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error at opening file", err)
		return
	}
	defer file.Close()

	_, err = file.Write(a)
	if err != nil {
		fmt.Println("Error Wrting FIle : ", err)
		return
	}
	fmt.Println("Data berhasil di buat Di : Storage/Task.json")
}
