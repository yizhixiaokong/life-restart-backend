package database

import (
	"context"
	"life-restart-backend/internal/config"
	"log"
	"time"

	"github.com/qiniu/qmgo"
)

var Client *qmgo.Client

func InitDatabase() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	Client, err = qmgo.NewClient(ctx, &qmgo.Config{Uri: config.AppConfig.DatabaseURI})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Database connected")
}

func GetCollection(collectionName string) *qmgo.Collection {
	return Client.Database(config.AppConfig.DatabaseName).Collection(collectionName)
}

func ShutdownHandler() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := Client.Close(ctx); err != nil {
		log.Fatalf("Failed to close database client: %v", err)
	}
	log.Println("Database client closed")
}
