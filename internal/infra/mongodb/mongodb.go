package mongodb

import (
	"context"
	"fmt"
	"log"
	"mercurio-web-scraping/internal/config"
	"mercurio-web-scraping/internal/domain/entities"
	"mercurio-web-scraping/internal/domain/link_handlers"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	DB      *mongo.Database
	context context.Context
	config  config.Config
}

func GetConnection(ctx context.Context, config config.Config) *Database {
	fmt.Println("Connecting with MongoDB...")
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.MongoURI))
	if err != nil {
		log.Fatalf("Error while connecting to mongo: %v\n", err)
	}
	fmt.Println("Connect with MongoDB")
	return &Database{DB: client.Database(config.MongoDBName), context: ctx, config: config}
}

func (db *Database) SeedDB() error {
	fmt.Println("Seeding MongoDB...")

	linkCollection := db.DB.Collection("link")

	linksToInsert := []interface{}{}

	for _, seedLink := range link_handlers.GetSeedLinks(db.config) {
		existLink := entities.Link{}
		err := linkCollection.FindOne(context.TODO(), bson.M{"url": seedLink.Url}).Decode(&existLink)
		if err != nil && err.Error() == "mongo: no documents in result" {
			seedLink.SetDefaultValues()
			linksToInsert = append(linksToInsert, seedLink)
		}
	}

	if len(linksToInsert) > 0 {
		_, err := linkCollection.InsertMany(context.TODO(), linksToInsert, options.MergeInsertManyOptions())
		if err != nil {
			log.Fatalf("Error while seeding mongo: %v\n", err)
		}
	}

	fmt.Println("Sown MongoDB")
	return nil
}
