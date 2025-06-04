package main

import (
	"bufio"
	"fmt"
	"os"

	"fundamentals/task"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Insert your Title: ")
	scanner.Scan()
	title := scanner.Text()

	err := task.AddTask(title)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
