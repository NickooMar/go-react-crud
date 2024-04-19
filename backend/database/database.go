package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDB es una estructura que contiene la información de la conexión a la base de datos
type MongoDB struct {
	client *mongo.Client
	db     *mongo.Database
}

func ConnectDatabase() (*MongoDB, error) {

	// Connection options
	connectionString := "mongodb://localhost:27017"
	clientOptions := options.Client().SetConnectTimeout(10 * time.Second)

	client, err := mongo.Connect(context.TODO(), clientOptions.ApplyURI(connectionString))
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	db := client.Database("simple_go_crud")

	// Creates a new mongoDB object
	mongoDB := &MongoDB{
		client: client,
		db:     db,
	}

	fmt.Println()

	return mongoDB, nil
}

func (m *MongoDB) Disconnect() {
	m.client.Disconnect(context.TODO())
}

func (m *MongoDB) GetCollection(collectionName string) *mongo.Collection {
	return m.db.Collection(collectionName)
}
