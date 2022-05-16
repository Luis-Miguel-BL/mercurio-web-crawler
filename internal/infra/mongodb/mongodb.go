package mongodb

import (
	"context"
	"log"
	"mercurio-web-scraping/internal/config"
	"mercurio-web-scraping/internal/domain/entities"
	"mercurio-web-scraping/internal/domain/seed"
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
	log.Println("Connecting with MongoDB...")
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.MongoURI))
	if err != nil {
		log.Fatalf("Error while connecting to mongo: %v\n", err)
	}
	log.Println("Connect with MongoDB")
	return &Database{DB: client.Database(config.MongoDBName), context: ctx, config: config}
}

func (db *Database) SeedDB() error {
	log.Println("Seeding MongoDB...")

	linkCollection := db.DB.Collection(entities.LinkCollectionName)
	notificationCollection := db.DB.Collection(entities.NotificationCollectionName)

	linksToInsert := []interface{}{}
	for _, seedLink := range seed.GetSeedLinks(db.config) {
		existLink := entities.Link{}
		err := linkCollection.FindOne(context.TODO(), bson.M{"url": seedLink.Url}).Decode(&existLink)
		if err != nil && err.Error() == "mongo: no documents in result" {
			seedLink.SetDefaultValues()
			linksToInsert = append(linksToInsert, seedLink)
		}
	}
	insertSeed(context.TODO(), linksToInsert, *linkCollection)

	notificationsToInsert := []interface{}{}
	for _, seedNotification := range seed.GetSeedNotifications(db.config) {
		existNotification := entities.Notification{}
		err := notificationCollection.FindOne(context.TODO(), bson.M{"contact": seedNotification.Contact, "channel": seedNotification.Channel, "harvest_target": seedNotification.HarvestTarget}).Decode(&existNotification)
		if err != nil && err.Error() == "mongo: no documents in result" {
			seedNotification.SetDefaultValues()
			notificationsToInsert = append(notificationsToInsert, seedNotification)
		}
	}
	insertSeed(context.TODO(), notificationsToInsert, *notificationCollection)

	log.Println("Sown MongoDB")
	return nil
}

func insertSeed(ctx context.Context, dataToInsert []interface{}, collection mongo.Collection) error {
	if len(dataToInsert) > 0 {
		_, err := collection.InsertMany(ctx, dataToInsert, options.MergeInsertManyOptions())
		if err != nil {
			log.Fatalf("Error while seeding mongo: %v\n", err)
			return err
		}
	}
	return nil
}
