package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
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
	var id int64
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

func genId() int64 {
	then := time.Date(2020, 1, 1, 23, 0, 0, 0, time.UTC)
	now := time.Now()
	duration := now.Sub(then)
	return duration.Milliseconds()
}
