package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type Connection struct {
	Ctx      context.Context
	client   *mongo.Client
	database *mongo.Database
}

func NewConnection(connectionString string, databaseName string) (*Connection, error) {
	ctx, _ := context.WithTimeout(context.Background(), 1000*time.Second)
	var err error
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		return nil, err
	}
	database := client.Database(databaseName)
	connection := Connection{Ctx: ctx, client: client, database: database}
	return &connection, nil
}

func (connection Connection) Disconnect() error {
	err := connection.client.Disconnect(connection.Ctx)
	return err
}

func (connection Connection) NewCollection(collectionName string) *mongo.Collection {
	collection := connection.database.Collection(collectionName)
	return collection
}
