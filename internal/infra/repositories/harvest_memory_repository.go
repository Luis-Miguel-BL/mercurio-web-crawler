package repositories

import (
	"context"
	"mercurio-web-crawler/internal/domain/entities"
)

type HarvestRepository struct {
	Harvests []entities.Harvest
}

func NewHarvestMemoryRepository() *HarvestRepository {
	return &HarvestRepository{Harvests: []entities.Harvest{}}
}

func (repo *HarvestRepository) Create(context context.Context, harvest entities.Harvest) (err error) {
	repo.Harvests = append(repo.Harvests, harvest)
	return nil
}
