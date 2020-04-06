package models

import (
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

const (
	COLLECTION = "tasks"
)

type Task struct {
	Id          string
	Status      string    `json:"status"`
	Text        string    `json:"text"`
	CreatedDate time.Time `json:"created_date"`
}

func (t *Task) Save() (interface{}, error) {
	c, err := NewDbClient()
	defer c.Client.Disconnect(c.ctx)

	if err != nil {
		return nil, err.(error)
	}
	err = c.Connect()
	if err != nil {
		return nil, err.(error)
	}

	tasks := c.GetCollection(COLLECTION)
	res, err := tasks.InsertOne(c.ctx, bson.M{"id": t.Id, "status": t.Status, "text": t.Text, "created_date": t.CreatedDate})
	if err != nil {
		return nil, err.(error)
	}

	fmt.Println("Salvando a task: " + t.Text)

	return res.InsertedID, nil
}

func (t *Task) Get() (Task, error) {
	emptyTask := Task{}
	var result bson.M
	c, err := NewDbClient()
	defer c.Client.Disconnect(c.ctx)

	if err != nil {
		return emptyTask, err.(error)
	}
	err = c.Connect()
	if err != nil {
		return emptyTask, err.(error)
	}

	tasks := c.GetCollection(COLLECTION)
	err = tasks.FindOne(c.ctx, bson.M{"id": t.Id}).Decode(&result)
	if err != nil {
		return emptyTask, err.(error)
	}
	resp, _ := json.Marshal(result)
	json.Unmarshal(resp, &emptyTask)

	fmt.Println("Retornando a task: " + t.Text)

	return emptyTask, nil
}

func (t *Task) Del() (bool, error) {
	c, err := NewDbClient()
	defer c.Client.Disconnect(c.ctx)

	if err != nil {
		return false, err.(error)
	}
	err = c.Connect()
	if err != nil {
		return false, err.(error)
	}

	tasks := c.GetCollection(COLLECTION)
	res, err := tasks.DeleteOne(c.ctx, bson.M{"id": t.Id})
	if err != nil {
		return false, err.(error)
	}

	fmt.Println("Retornando a task: ", res.DeletedCount)

	return true, nil
}

func GetAll() ([]Task, error) {
	tasks := []Task{}
	fmt.Println("Pegando todas as tasks")
	return tasks, nil
}
