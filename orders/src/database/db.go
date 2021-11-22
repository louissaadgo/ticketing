package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	DB *mongo.Collection
)

func Connect() {
	uri := "mongodb://orders-mongo-srv:27017/"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		//Handle the error and restart the pod later
		fmt.Println(err)
	}
	DB = client.Database("orders").Collection("orders")
}
