package service

import (
	"mercurio-web-scraping/internal/domain/contract"
)

type Service struct {
	LinkService    LinkService
	HarvestService HarvestService
}

func GetServices(linkRepo contract.LinkRepository, harvestRepo contract.HarvestRepository) *Service {
	return &Service{LinkService: *NewLinkService(linkRepo), HarvestService: *NewHarvestService(harvestRepo)}
}
