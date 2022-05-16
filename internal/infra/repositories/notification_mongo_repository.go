package repositories

import (
	"context"
	"mercurio-web-scraping/internal/domain/entities"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type NotificationMongoRepository struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewNotificationMongoRepository(mongo *mongo.Database) *NotificationMongoRepository {
	return &NotificationMongoRepository{db: mongo, collection: mongo.Collection("notification")}
}
func (repo *NotificationMongoRepository) FindByTarget(ctx context.Context, target entities.HarvestType) (notifications []entities.Notification, err error) {
	cursor, err := repo.collection.Find(ctx, bson.M{"harvest_target": target})
	if err != nil {
		return notifications, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var notification entities.Notification
		err := cursor.Decode(&notification)
		if err != nil {
			return notifications, err
		}

		notifications = append(notifications, notification)

	}

	return notifications, nil
}
