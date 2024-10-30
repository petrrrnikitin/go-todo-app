package tasks

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

const fileName = "tasks.json"

type Task struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func (t *Task) MarkAsCompleted() {
	t.Completed = true
}

func New(id int, title string) (Task, error) {
	if title == "" {
		return Task{}, errors.New("title is required")
	}
	return Task{
		Id:        id,
		Title:     title,
		Completed: false,
	}, nil
}

func DisplayList(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("Empty list")
		return
	}

	for i, task := range tasks {
		fmt.Printf("%v) %v, completed: %t\n", i+1, task.Title, task.Completed)
	}
}

func LoadList() ([]Task, error) {
	content, err := os.ReadFile(fileName)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return []Task{}, nil
		}
		return nil, err
	}

	var tasks []Task
	err = json.Unmarshal(content, &tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func SaveToFile(tasks []Task) error {
	content, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, content, 0644)
}
