package service

import (
	"context"
	"mercurio-web-crawler/internal/domain/contract"
	"mercurio-web-crawler/internal/domain/entities"
)

type HarvestService struct {
	repo contract.HarvestRepository
}

func NewHarvestService(repo contract.HarvestRepository) *HarvestService {
	return &HarvestService{
		repo: repo,
	}
}

func (s *HarvestService) Create(context context.Context, harvest entities.Harvest) (err error) {
	harvest.SetDefaultValues()
	return s.repo.Create(context, harvest)
}
