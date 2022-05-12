package repositories

import (
	"context"
	"mercurio-web-crawler/internal/domain/entities"
)

type HarvestMemoryRepository struct {
	Harvests []entities.Harvest
}

func NewHarvestMemoryRepository() *HarvestMemoryRepository {
	return &HarvestMemoryRepository{Harvests: []entities.Harvest{}}
}

func (repo *HarvestMemoryRepository) Create(context context.Context, harvest entities.Harvest) (err error) {
	repo.Harvests = append(repo.Harvests, harvest)
	return nil
}
