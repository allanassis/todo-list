package models

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

const (
	COLLECTION = "tasks"
)

type Task struct {
	Id          string
	Status      string `json:"status"`
	Text        string `json:"text"`
	CreatedDate time.Time
}

func (t *Task) Save() (interface{}, error) {
	client, err := NewDbClient()
	if err != nil {
		return nil, err.(error)
	}
	err = client.Connect()
	if err != nil {
		return nil, err.(error)
	}

	tasks := client.GetCollection(COLLECTION)
	res, err := tasks.InsertOne(client.ctx, bson.M{"id": t.Id, "status": t.Status, "text": t.Text, "created_date": t.CreatedDate})
	if err != nil {
		return nil, err.(error)
	}

	fmt.Println("Salvando a task: " + t.Text)

	return res.InsertedID, nil
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
