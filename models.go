package main

import (
	"github.com/google/uuid"
)

type Task struct {
	ID        string `json:"id"`
	Text      string `json:"text"`
	Completed bool   `json:"completed"`
}

func addTask(text string, tasks *[]Task) uuid.UUID {
	id := uuid.New()
	completed := false

	var new_task Task = Task{id.String(), text, completed}

	*tasks = append(*tasks, new_task)

	return id
}
