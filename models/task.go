package models

import (
	"fmt"
	"time"
)

type Task struct {
	id          string
	status      string
	text        string
	createdDate time.Time
}

func (t *Task) Save() error {
	fmt.Println("Salvando a task: %s", t.text)
	return nil
}

func (t *Task) Del() error {
	fmt.Println("Deletando a task: %s", t.text)
	return nil
}

func GetAll() ([]Task, error) {
	tasks := []Task{}
	fmt.Println("Pegando todas as tasks")
	return tasks, nil
}
