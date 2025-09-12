package main

import "fmt"

type TodoItem struct {
	task      string
	id        int
	completed bool
}

func main() {
	const OPTION_BEGIN uint = 1
	const OPTION_END uint = 6

	var option uint
	var tasks []TodoItem

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
			tasks[taskIdx].completed = true
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
			if task.completed {
				status = "completed"
			} else {
				status = "pending"
			}
			fmt.Printf("Task: %v\tTask ID: %v\tStatus: %v\n", task.task, task.id, status)
		}
	} else {
		fmt.Println("No tasks exist")
	}
}

func addNewTask(tasksPointer *[]TodoItem) {
	var newTask string = getTask()
	var todoItem = TodoItem{
		task:      newTask,
		id:        len(*tasksPointer),
		completed: false,
	}
	*tasksPointer = append(*tasksPointer, todoItem)
}

func editTask(tasksPointer *[]TodoItem, taskIdx int) {
	var newTask string = getTask()
	(*tasksPointer)[taskIdx].task = newTask
}

func deleteTask(tasksPointer *[]TodoItem, taskIdx int) {
	*tasksPointer = append((*tasksPointer)[:taskIdx], (*tasksPointer)[taskIdx+1:]...)
}
