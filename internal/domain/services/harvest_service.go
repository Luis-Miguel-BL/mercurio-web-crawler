package service

import (
	"context"
	"mercurio-web-scraping/internal/domain/contract"
	"mercurio-web-scraping/internal/domain/entities"
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
	err = harvest.Validate()
	if err != nil {
		return err
	}
	return s.repo.Create(context, harvest)
}

func (s *HarvestService) FindByPageLink(context context.Context, pageLink string) (harvest entities.Harvest, err error) {
	return s.repo.FindByPageLink(context, pageLink)
}
func (s *HarvestService) Update(context context.Context, harvest entities.Harvest) (err error) {
	err = harvest.Validate()
	if err != nil {
		return err
	}
	return s.repo.Update(context, harvest)
}
