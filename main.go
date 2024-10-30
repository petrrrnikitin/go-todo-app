package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"todoApp/tasks"
)

func main() {
	taskList := tasks.LoadList()
	for {
		menu()
		input := getUserChoice("Your choice: ")

		switch input {
		case 1:
			tasks.DisplayList(taskList)
		case 2:
			taskTitle := getUserInput("Task text:")
			currentIndex := len(taskList)
			task, _ := tasks.New(currentIndex, taskTitle)
			taskList[currentIndex] = task
		case 3:
			tasks.DisplayList(taskList)
			id := getUserChoice("Task ID to remove")
			delete(taskList, id-1)
		case 4:
			id := getUserChoice("Task ID to mark as completed")
			task := taskList[id-1]
			task.MarkAsCompleted()
		case 5:
			err := tasks.SaveToFile(taskList)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println("See u soon")
			return
		}
	}
}

func getUserChoice(prompt string) int {
	var choice int
	fmt.Printf("%v ", prompt)
	fmt.Scan(&choice)
	return choice
}

func menu() {
	fmt.Println("Menu")
	fmt.Println("1. List tasks")
	fmt.Println("2. Add new task")
	fmt.Println("3. Remove task")
	fmt.Println("4. Mark as completed")
	fmt.Println("5. Exit")

}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
