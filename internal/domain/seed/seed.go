package seed

import (
	"mercurio-web-scraping/internal/application/notification"
	"mercurio-web-scraping/internal/config"
	"mercurio-web-scraping/internal/domain/entities"
)

func GetSeedLinks(config config.Config) []entities.Link {
	return []entities.Link{
		{Url: config.ZapImoveisURL, Slug: config.ZapImoveisSlug, Origin: "ZapImoveis", Description: "Novas Casas ZapImoveis", TimeoutInSeconds: 60, Active: true},
	}
}

func GetSeedNotifications(config config.Config) []entities.Notification {
	return []entities.Notification{
		{Channel: notification.NotificationChannelTwitter, Contact: "@Luis_MBL", HarvestTarget: entities.HarvestBuilding},
	}
}
