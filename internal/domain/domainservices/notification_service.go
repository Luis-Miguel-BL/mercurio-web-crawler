package domainservices

import (
	"context"
	"mercurio-web-scraping/internal/application/notification"
	"mercurio-web-scraping/internal/domain/contract"
	"mercurio-web-scraping/internal/domain/entities"
	"mercurio-web-scraping/internal/domain/notifications"
)

type NotificationService struct {
	repo                contract.NotificationRepository
	notificationService notification.NotificationService
}

func NewNotificationService(repo contract.NotificationRepository, notificationService notification.NotificationService) *NotificationService {
	return &NotificationService{
		repo:                repo,
		notificationService: notificationService,
	}
}

func (s NotificationService) FindAndNotifyByTargets(ctx context.Context, harvest entities.Harvest) (err error) {
	targets, err := s.repo.FindByTarget(ctx, harvest.HarvestType)
	if err != nil {
		return err
	}
	for _, target := range targets {
		notification := notifications.BuildNewHarvestNotification(notifications.NewHarvestNotificationData{
			Harvest:     harvest,
			Destination: target.Contact,
		})
		includeChannel := false
		for _, channel := range notification.Channels {
			if channel == target.Channel {
				includeChannel = true
			}
		}

		if includeChannel {
			s.notificationService.Notificate(*notification)
		}
	}

	return nil
}
