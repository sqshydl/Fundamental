package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"fundamentals/task"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n=== TODO APP ===")
		fmt.Println("1. Tambah Task")
		fmt.Println("2. Lihat Semua Task")
		fmt.Println("3. Hapus Task")
		fmt.Println("4. Keluar")
		fmt.Print("Pilih menu: ")
		scanner.Scan()
		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "1":
			fmt.Print("Masukkan judul task: ")
			scanner.Scan()
			title := scanner.Text()
			err := task.AddTask(title)
			if err != nil {
				fmt.Println("Gagal menambahkan task:", err)
			} else {
				fmt.Println("Task berhasil ditambahkan.")
			}
		case "2":
			task.ReadTasks()
		case "3":
			fmt.Print("Masukkan ID task yang ingin dihapus: ")
			scanner.Scan()
			idStr := scanner.Text()
			id, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Println("ID harus berupa angka.")
				continue
			}
			err = task.DeleteTask(id)
			if err != nil {
				fmt.Println("Gagal menghapus task:", err)
			} else {
				fmt.Println("Task berhasil dihapus.")
			}
		case "4":
			fmt.Println("Keluar dari aplikasi. 👋")
			return
		default:
			fmt.Println("Menu tidak dikenal.")
		}
	}
}
