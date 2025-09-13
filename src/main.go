package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type TodoItem struct {
	Task      string `json:"task"`
	Id        int    `json:"id"`
	Completed bool   `json:"completed"`
}

func main() {
	const OPTION_BEGIN uint = 1
	const OPTION_END uint = 6
	const FILE_PERM = 0644
	const FILE_PATH string = "todos.json"

	var option uint
	var tasks []TodoItem

	if !fileExists(FILE_PATH) {
		f, err := os.Create(FILE_PATH)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		defer f.Close()
	} else {
		jsonData, err := os.ReadFile(FILE_PATH)
		if err != nil {
			fmt.Println("Error occured while reading JSON file")
			os.Exit(1)
		}
		fmt.Printf("Read JSON data: %v\n", string(jsonData))
		unmarshalErr := json.Unmarshal(jsonData, &tasks)
		if unmarshalErr != nil {
			fmt.Println("Error occured while parsing JSON data")
			os.Exit(1)
		}
		fmt.Printf("Tasks after population: %v\n", tasks)
	}

	for {
		displayOptions()
		fmt.Printf("Choose option: ")
		fmt.Scan(&option)
		if option < OPTION_BEGIN || option > OPTION_END {
			fmt.Println("Invalid option!")
			continue
		}

		if option == 6 {
			break
		}

		switch option {
		case 1:
			viewTasks(tasks)
		case 2:
			addNewTask(&tasks)
		case 3:
			taskIdx, err := getTaskIndex(tasks)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			tasks[taskIdx].Completed = true
		case 4:
			taskIdx, err := getTaskIndex(tasks)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			editTask(&tasks, taskIdx)
		case 5:
			taskIdx, err := getTaskIndex(tasks)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			deleteTask(&tasks, taskIdx)
		}
	}

	jsonData, err := json.Marshal(tasks)
	if err != nil {
		fmt.Println("Error while encoding JSON data")
		os.Exit(1)
	}
	fmt.Printf("Encoded JSON data: %v\n", string(jsonData))
	os.WriteFile(FILE_PATH, jsonData, FILE_PERM)
}

func displayOptions() {
	fmt.Println("1. View tasks")
	fmt.Println("2. Add new task")
	fmt.Println("3. Complete task")
	fmt.Println("4. Edit task")
	fmt.Println("5. Delete task")
	fmt.Println("6. Exit")
}

func viewTasks(tasks []TodoItem) {
	if len(tasks) != 0 {
		for _, task := range tasks {
			var status string
			if task.Completed {
				status = "completed"
			} else {
				status = "pending"
			}
			fmt.Printf("Task: %v\tTask ID: %v\tStatus: %v\n", task.Task, task.Id, status)
		}
	} else {
		fmt.Println("No tasks exist")
	}
}

func addNewTask(tasksPointer *[]TodoItem) {
	var newTask string = getTask()
	var todoItem = TodoItem{
		Task:      newTask,
		Id:        len(*tasksPointer),
		Completed: false,
	}
	*tasksPointer = append(*tasksPointer, todoItem)
}

func editTask(tasksPointer *[]TodoItem, taskIdx int) {
	var newTask string = getTask()
	(*tasksPointer)[taskIdx].Task = newTask
}

func deleteTask(tasksPointer *[]TodoItem, taskIdx int) {
	*tasksPointer = append((*tasksPointer)[:taskIdx], (*tasksPointer)[taskIdx+1:]...)
}
