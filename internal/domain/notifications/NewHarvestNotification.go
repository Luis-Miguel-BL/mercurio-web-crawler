package notifications

import (
	"fmt"
	"mercurio-web-scraping/internal/application/notification"
	"mercurio-web-scraping/internal/domain/entities"
)

type NewHarvestNotificationData struct {
	Harvest     entities.Harvest
	Destination string
}

func BuildNewHarvestNotification(data NewHarvestNotificationData) *notification.Notification {
	translateHarvestType := map[entities.HarvestType]string{
		entities.HarvestBuilding: "moradia",
	}

	message := fmt.Sprintf(" Novo(a) %s encontrada no site %s", translateHarvestType[data.Harvest.HarvestType], data.Harvest.PageLink)

	return &notification.Notification{
		Channels:    []notification.NotificationChannel{notification.NotificationChannelTwitter},
		Destination: data.Destination,
		Message:     message,
	}
}
