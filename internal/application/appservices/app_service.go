package appservices

import (
	"mercurio-web-scraping/internal/application/notification"
	"mercurio-web-scraping/internal/config"
)

type Service struct {
	Notification *notification.NotificationService
}

func GetService(config config.Config) *Service {
	return &Service{
		Notification: notification.NewNotificationService(config),
	}
}
