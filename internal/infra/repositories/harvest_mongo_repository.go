package repositories

import (
	"context"
	"mercurio-web-scraping/internal/domain/entities"

	"go.mongodb.org/mongo-driver/mongo"
)

type HarvestMongoRepository struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewHarvestMongoRepository(mongo *mongo.Database) *HarvestMongoRepository {
	return &HarvestMongoRepository{db: mongo, collection: mongo.Collection("harvest")}
}

func (repo *HarvestMongoRepository) Create(context context.Context, harvest entities.Harvest) (err error) {
	_, err = repo.collection.InsertOne(context, harvest)
	return err
}
