package repositories

import (
	"context"
	"mercurio-web-scraping/internal/domain/entities"

	"go.mongodb.org/mongo-driver/bson"
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
func (repo *HarvestMongoRepository) FindByPageLink(context context.Context, pageLink string) (harvest entities.Harvest, err error) {
	err = repo.collection.FindOne(context, bson.M{"page_link": pageLink}).Decode(&harvest)
	return harvest, err
}
func (repo *HarvestMongoRepository) Update(ctx context.Context, harvest entities.Harvest) (err error) {
	_, err = repo.collection.UpdateOne(ctx, bson.M{"_id": harvest.ID}, buildHarvestToUpdate(harvest))
	return err
}

func buildHarvestToUpdate(harvest entities.Harvest) bson.M {
	return bson.M{
		"$set": bson.M{
			"raw_data":       harvest.RawData,
			"page_link":      harvest.PageLink,
			"info":           harvest.Info,
			"disappeared_at": harvest.DisappearedAt,
		},
	}
}
