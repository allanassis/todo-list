package models

import (
	"fmt"
	"time"
)

type Task struct {
	Id          string
	Status      string
	Text        string
	CreatedDate time.Time
}

func (t *Task) Save() error {
	fmt.Println("Salvando a task: %s", t.Text)
	return nil
}

func (t *Task) Del() error {
	fmt.Println("Deletando a task: %s", t.Text)
	return nil
}

func GetAll() ([]Task, error) {
	tasks := []Task{}
	fmt.Println("Pegando todas as tasks")
	return tasks, nil
}
