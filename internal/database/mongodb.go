package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

type MongoDB struct {
	Client   *mongo.Client
	Database *mongo.Database
}

func NewMongoDBWithOptions(uri, dbName string, otps *options.ClientOptions) (*MongoDB, func(), error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	otps.ApplyURI(uri)

	//Connect to MonogDB
	client, err := mongo.Connect(otps)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	// Verify connection
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		client.Disconnect(ctx)
		return nil, nil, fmt.Errorf("failed to ping MongoDB: %w", err)

	}

	log.Println("Connected to MongoDB successfully")

	db := client.Database(dbName)

	MongoDB := &MongoDB{
		Client:   client,
		Database: db,
	}

	cleanup := func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := client.Disconnect(ctx); err != nil {
			log.Printf("Error disconnecting from MongoDB: %v", err)
		} else {
			log.Println("Disconnected from MongoDB")
		}

	}
	return MongoDB, cleanup, nil
}
