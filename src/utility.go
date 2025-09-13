package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func getTask() string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Enter new task: ")
	task, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading from stdin!")
		os.Exit(1)
	}

	return strings.TrimSpace(task)
}

func getTaskIndex(tasks []TodoItem) (int, error) {
	var id int
	var idx int = -1

	fmt.Printf("Enter task ID: ")
	fmt.Scan(&id)
	for index, task := range tasks {
		if task.Id == id {
			idx = index
			break
		}
	}
	if idx == -1 {
		return idx, errors.New("Not a valid task ID")
	}

	return idx, nil
}
