package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	MONGO_URL = "mongodb://localhost:27017"
	DATA_BASE = "todo-list"
)

type DbClient struct {
	Client *mongo.Client
	ctx    context.Context
}

func NewDbClient() (DbClient, interface{}) {

	ctx, err := context.WithTimeout(context.Background(), 10*time.Second)
	if err != nil {
		return DbClient{}, err
	}

	client, err2 := mongo.NewClient(options.Client().ApplyURI(MONGO_URL))
	if err2 != nil {
		return DbClient{}, err2
	}

	return DbClient{
		Client: client,
		ctx:    ctx,
	}, nil

}

func (db *DbClient) Connect() error {

	err := db.Client.Connect(db.ctx)
	if err != nil {
		return err
	}

	if err := db.Client.Ping(db.ctx, readpref.Primary()); err != nil {
		return err
	}

	return nil
}

func (db *DbClient) GetCollection(collection string) *mongo.Collection {
	return db.Client.Database(DATA_BASE).Collection(collection)
}

// func Bla() {

// 	collection := client.Database("testing").Collection("numbers")
// 	res, _ := collection.InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})
// 	fmt.Println(res.InsertedID)
// }
