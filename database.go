package main

import (
	"context"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var once sync.Once
var client *mongo.Client
var err error

func Dabase() (*mongo.Client, error) {
	once.Do(func() {
		// TOOD: move the uri to somewhere else
		client, err = mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://fitz:fitz@localhost:27017/?authSource=admin"))

		if err != nil {
			return
		}

		if err = client.Ping(context.Background(), nil); err != nil {
			return
		}
	})

	if err != nil {
		return nil, err
	}
	return client, nil
}
