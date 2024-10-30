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

func DisplayList(c map[int]Task) {
	if len(c) == 0 {
		fmt.Println("Empty list")
	}

	for key, value := range c {
		key += 1
		fmt.Printf("%v) : %v, completed: %t \n", key, value.Title, value.Completed)
	}
}

func LoadList() map[int]Task {
	content, err := os.ReadFile(fileName)
	if err != nil {
		return make(map[int]Task)
	}

	var collection map[int]Task
	err = json.Unmarshal(content, &collection)

	if err != nil {
		return make(map[int]Task)
	}

	return collection
}

func SaveToFile(collection map[int]Task) error {
	list, err := json.Marshal(collection)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, list, 0644)
}
