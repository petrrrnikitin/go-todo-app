package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"todoApp/tasks"
)

func main() {
	taskList, err := tasks.LoadList()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	for {
		menu()
		input := getUserChoice("Your choice: ")

		switch input {
		case 1:
			tasks.DisplayList(taskList)
		case 2:
			addNewTask(&taskList)
		case 3:
			removeTask(&taskList)
		case 4:
			markTaskAsCompleted(&taskList)
		case 5:
			saveAndExit(taskList)
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}

func addNewTask(taskList *[]tasks.Task) {
	taskTitle := getUserInput("Task text:")
	task, err := tasks.New(len(*taskList)+1, taskTitle)
	if err != nil {
		fmt.Println(err)
		return
	}
	*taskList = append(*taskList, task)
}

func removeTask(taskList *[]tasks.Task) {
	tasks.DisplayList(*taskList)
	id := getUserChoice("Task ID to remove: ")
	if id < 1 || id > len(*taskList) {
		fmt.Println("Invalid ID. Task does not exist.")
		return
	}
	*taskList = append((*taskList)[:id-1], (*taskList)[id:]...)
}

func markTaskAsCompleted(taskList *[]tasks.Task) {
	id := getUserChoice("Task ID to mark as completed: ")
	if id < 1 || id > len(*taskList) {
		fmt.Println("Invalid ID. Task does not exist.")
		return
	}
	(*taskList)[id-1].MarkAsCompleted()
}

func saveAndExit(taskList []tasks.Task) {
	err := tasks.SaveToFile(taskList)
	if err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}
	fmt.Println("See you soon")
	os.Exit(0)
}

func menu() {
	fmt.Println("Menu")
	fmt.Println("1. List tasks")
	fmt.Println("2. Add new task")
	fmt.Println("3. Remove task")
	fmt.Println("4. Mark as completed")
	fmt.Println("5. Exit")
}

func getUserChoice(prompt string) int {
	var choice int
	fmt.Printf("%v ", prompt)
	_, err := fmt.Scan(&choice)
	if err != nil {
		fmt.Println("Invalid input. Please enter a number.")
		return getUserChoice(prompt)
	}
	return choice
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}
