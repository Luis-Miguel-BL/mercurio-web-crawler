package mongodb

import (
	"context"
	"log"
	"mercurio-web-crawler/internal/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connection(ctx context.Context, config config.Config) *mongo.Database {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.MongoURI))
	if err != nil {
		log.Fatalf("Error while connecting to mongo: %v\n", err)
	}
	return client.Database(config.MongoDBName)
}
