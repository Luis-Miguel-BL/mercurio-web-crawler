package repositories

import (
	"context"
	"mercurio-web-scraping/internal/domain/entities"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type LinkMongoRepository struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewLinkMongoRepository(mongo *mongo.Database) (repo *LinkMongoRepository) {

	return &LinkMongoRepository{
		db: mongo, collection: mongo.Collection(entities.LinkCollectionName),
	}
}

func (repo *LinkMongoRepository) GetByUUID(ctx context.Context, LinkUUID string) (link entities.Link, err error) {
	err = repo.collection.FindOne(ctx, bson.M{"uuid": LinkUUID}).Decode(&link)
	return link, err

}

func (repo *LinkMongoRepository) FindAvailableToVisits(ctx context.Context) (links []entities.Link, err error) {
	cursor, err := repo.collection.Find(ctx, bson.M{"active": true})
	if err != nil {
		return links, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(context.TODO()) {
		var link entities.Link
		err := cursor.Decode(&link)
		if err != nil {
			return links, err
		}

		if link.LastVisit.Unix()+link.TimeoutInSeconds < time.Now().Unix() {
			links = append(links, link)
		}

	}

	return links, nil
}

func (repo *LinkMongoRepository) Update(ctx context.Context, link entities.Link) (err error) {
	_, err = repo.collection.UpdateOne(ctx, bson.M{"_id": link.ID}, buildLinkToUpdate(link))
	return err
}

func buildLinkToUpdate(link entities.Link) bson.M {
	return bson.M{
		"$set": bson.M{
			"url":                   link.Url,
			"slug":                  link.Slug,
			"origin":                link.Origin,
			"description":           link.Description,
			"last_visit":            link.LastVisit,
			"timeout_in_seconds":    link.TimeoutInSeconds,
			"active":                link.Active,
			"total_visits":          link.TotalVisits,
			"total_error_in_visits": link.TotalErrorInVisits,
		},
	}
}
