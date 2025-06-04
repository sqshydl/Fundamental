package task

import (
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

func AddTask(title string) error {
	os.MkdirAll("storage", os.ModePerm)
	tasks := loadTasks()
	ai := autoInc{id: len(tasks)}

	newTask := Task{Id: ai.ID(), Title: title, Completed: false}
	tasks = append(tasks, newTask)

	err := saveTasks(tasks)
	if err != nil {
		return err
	}

	fmt.Println("Task berhasil ditambahkan ke storage/Task.json")
	return nil
}

func loadTasks() []Task {
	filename := "storage/Task.json"
	var tasks []Task

	data, err := os.ReadFile(filename)
	if err != nil {
		// Kalau file belum ada, kembalikan slice kosong
		return []Task{}
	}

	err = json.Unmarshal(data, &tasks)
	if err != nil {
		fmt.Println("Gagal membaca file JSON:", err)
		return []Task{}
	}

	return tasks
}

func saveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile("storage/Task.json", data, 0644)
}

func readTasks() {
	tasks := loadTasks()
	if len(tasks) == 0 {
		fmt.Println("Belum ada task.")
		return
	}

	for _, task := range tasks {
		status := "Belum selesai"
		if task.Completed {
			status = "Selesai"
		}
		fmt.Printf("ID: %d | Title: %s | Status: %s\n", task.Id, task.Title, status)
	}
}

func deleteTask() {

}
