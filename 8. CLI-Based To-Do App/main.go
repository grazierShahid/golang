package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Display() {
	fmt.Println("CLI-Based To-Do App Task")
	fmt.Println("1. Add a task")
	fmt.Println("2. List tasks")
	fmt.Println("3. Mark a Task as Completed")
	fmt.Println("4. Delete a Task")
	fmt.Println("5. Exit")
}

func getChoice(reader *bufio.Reader) (int, error) {
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	return strconv.Atoi(input)
}

func getTask(reader *bufio.Reader) string {
	task, _ := reader.ReadString('\n')
	return strings.TrimSpace(task)
}

func markTaskCompleted(tasks []string, taskId int) ([]string, bool) {
	if taskId > 0 && taskId <= len(tasks) {
		tasks[taskId-1] = "[+]" + tasks[taskId-1]
		fmt.Println("Task marked as completed.")

		return tasks, true
	}

	fmt.Println("Invalid task number.")
	return tasks, false
}

func deleteTask(tasks []string, taskId int) ([]string, bool) {
	if taskId > 0 && taskId <= len(tasks) {
		tasks = append(tasks[:taskId-1], tasks[taskId:]...)
		fmt.Print("Task deleted!\n")

		return tasks, true
	}

	fmt.Println("Invalid task number.")
	return tasks, false
}

func main() {
	var tasks []string
	reader := bufio.NewReader(os.Stdin)

	for {
		Display()
		fmt.Printf("\nEnter your choice: ")

		choice, err := getChoice(reader)
		if err != nil {
			fmt.Println("Invalid input, please enter a number.")
		}

		switch choice {
		case 1:
			// Add task
			fmt.Println("Enter the task: ")

			// add this task to the task list
			tasks = append(tasks, getTask(reader))
			fmt.Println("Task added.")

		case 2:
			// Show task list
			if len(tasks) == 0 {
				fmt.Println("No task is availabe.")
			} else {
				fmt.Println("Your task:")
				for index, task := range tasks {
					fmt.Printf("%d - %s\n", index+1, task)
				}
			}

		case 3:
			// Mark task completed
			if len(tasks) == 0 {
				fmt.Println("No task is availabe.")
			}

			fmt.Print("Enter task num to complete: ")

			taskId, err := getChoice(reader)
			if err == nil {
				tasks, _ = markTaskCompleted(tasks, taskId)
			} else {
				fmt.Println("Invalid input, please enter a number.")
			}

		case 4:
			// Delete task
			if len(tasks) == 0 {
				fmt.Println("No task is availabe.")
			}
			fmt.Print("Enter task num to delete: ")

			taskId, err := getChoice(reader)

			if err == nil {
				tasks, _ = deleteTask(tasks, taskId)
			} else {
				fmt.Println("Invalid input, please enter a number.")
			}

		case 5:
			fmt.Println("Exiting. Goodbye....")
			os.Exit(1)
		default:
			fmt.Println("Invalide choice, try again.")

		}
	}
}
