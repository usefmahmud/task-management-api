package main

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID        string `json:"id"`
	Text      string `json:"text"`
	Completed bool   `json:"completed"`
	Time      string `json:"time"`
}

func addTask(text string, tasks *[]Task) uuid.UUID {
	id := uuid.New()
	completed := false

	var new_task Task = Task{
		ID:        id.String(),
		Text:      text,
		Completed: completed,
		Time:      time.Now().String(),
	}

	*tasks = append(*tasks, new_task)

	return id
}

func getTask(id string, tasks *[]Task) (err error, task *Task) {
	for _, task := range *tasks {
		if task.ID == id {
			return nil, &task
		}
	}

	return errors.New("Task is not exist"), nil
}

func completeTask(id string, tasks *[]Task) (err error, task *Task) {
	for i := range *tasks {
		if (*tasks)[i].ID == id {
			(*tasks)[i].Completed = true
			return nil, &(*tasks)[i]
		}
	}

	return errors.New("Task is not exist"), nil
}

func removeTask(id string, tasks *[]Task) (err error) {
	for i := range *tasks {
		if (*tasks)[i].ID == id {
			*tasks = append((*tasks)[:i], (*tasks)[i+1:]...)
			return nil
		}
	}

	return errors.New("Task is not exist")
}
