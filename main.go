package main

import (
	"bufio"
	"fmt"
	"os"

	"fundamentals/task" // change this to match your go.mod module name
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
