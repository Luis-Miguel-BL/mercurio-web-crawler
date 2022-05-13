package mongodb

import (
	"context"
	"fmt"
	"log"
	"mercurio-web-scraping/internal/config"
	"mercurio-web-scraping/internal/domain/entities"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	DB      *mongo.Database
	context context.Context
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
	return &Database{DB: client.Database(config.MongoDBName), context: ctx}
}

func (db *Database) SeedDB() error {
	fmt.Println("Seeding MongoDB...")
	linkCollection := db.DB.Collection("link")

	linksToInsert := []interface{}{}

	for _, seedLink := range SeedLinks {
		existLink := entities.Link{}
		err := linkCollection.FindOne(context.TODO(), bson.M{"url": seedLink.Url}).Decode(&existLink)
		if err != nil && err.Error() == "mongo: no documents in result" {
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

var SeedLinks []entities.Link = []entities.Link{
	{Url: "https://google.com", Slug: "teste-google", Origin: "google", Description: "teste google", TimeoutInSeconds: 5, HarvestType: entities.HarvestBuilding, Active: true},
	{Url: "https://facebook.com", Slug: "teste-facebook", Origin: "facebook", Description: "teste facebook", TimeoutInSeconds: 5, HarvestType: entities.HarvestBuilding, Active: true},
}
