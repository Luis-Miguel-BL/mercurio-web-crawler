package domainservices

import (
	"mercurio-web-scraping/internal/application/appservices"
	"mercurio-web-scraping/internal/domain/contract"
)

type Service struct {
	LinkService         LinkService
	HarvestService      HarvestService
	NotificationService NotificationService
}

func GetServices(appServices appservices.Service, linkRepo contract.LinkRepository, harvestRepo contract.HarvestRepository, notificationRepo contract.NotificationRepository) *Service {
	return &Service{LinkService: *NewLinkService(linkRepo), HarvestService: *NewHarvestService(harvestRepo), NotificationService: *NewNotificationService(notificationRepo, *appServices.Notification)}
}
